import { Component, OnInit, OnDestroy, EventEmitter } from '@angular/core';
import { Store } from '@ngrx/store';
import { Subject, combineLatest } from 'rxjs';
import { filter, takeUntil, pluck } from 'rxjs/operators';
import { identity, isNil } from 'lodash/fp';
import { NgrxStateAtom } from 'app/ngrx.reducers';
import { EntityStatus } from 'app/entities/entities';
import { LayoutFacadeService, Sidebar } from 'app/entities/layout/layout.facade';
import { routeParams } from 'app/route.selectors';

import {
  DeleteDataBagItem,
  GetDataBagItems
} from 'app/entities/data-bags/data-bag-details.actions';
import { DataBagItems, DataBagsItemDetails } from 'app/entities/data-bags/data-bags.model';
import { getAllStatus, dataBagItemList, deleteStatus } from 'app/entities/data-bags/data-bag-details.selector';
import { GetDataBagItemDetails } from 'app/entities/data-bags/data-bag-item-details.actions';
import { dataBagItemDetailsFromRoute, getStatus  } from 'app/entities/data-bags/data-bag-item-details.selector';

export type DataBagsDetailsTab = 'details';

@Component({
  selector: 'app-data-bags-details',
  templateUrl: './data-bags-details.component.html',
  styleUrls: ['./data-bags-details.component.scss']
})
export class DataBagsDetailsComponent implements OnInit, OnDestroy {
  private isDestroyed = new Subject<boolean>();
  public dataBagItems: DataBagItems[];
  public dataBagSearch: DataBagItems[];
  public dataBagItemDetails: DataBagsItemDetails;
  public serverId: string;
  public orgId: string;
  public dataBagName: string;
  public dataBagItemName: string;
  public itemDataJson: string;
  public tabValue: DataBagsDetailsTab = 'details';
  public dataBagsDetailsLoading = true;
  public selectedItem: string;
  public dataBagsItemDetailsLoading = false;
  public selectedItemDetails: object;
  public activeClassName: string;
  public searching = false;
  public searchValue = '';
  public current_page = 1;
  public per_page = 9;
  public total: number;
  public dataBagItemToDelete: DataBagItems;
  public deleteModalVisible = false;
  public openDataBagModal = new EventEmitter<void>();
  public openEditDataBagItemModal = new EventEmitter<void>();
  public openDataBagItemModal = new EventEmitter<void>();

  constructor(
    private store: Store<NgrxStateAtom>,
    private layoutFacade: LayoutFacadeService
  ) { }

  ngOnInit() {
    this.layoutFacade.showSidebar(Sidebar.Infrastructure);

    combineLatest([
      this.store.select(routeParams).pipe(pluck('id'), filter(identity)),
      this.store.select(routeParams).pipe(pluck('org-id'), filter(identity)),
      this.store.select(routeParams).pipe(pluck('name'), filter(identity))
    ]).pipe(
      takeUntil(this.isDestroyed)
    ).subscribe(([server_id, org_id, dataBag_name]) => {
      this.serverId = server_id;
      this.orgId = org_id;
      this.dataBagName = dataBag_name;
      this.getDataBagItemsData();
    });

    combineLatest([
      this.store.select(getAllStatus),
      this.store.select(dataBagItemList)
    ]).pipe(
      filter(([getDataBagItemsSt, _dataBagItemsState]) =>
      getDataBagItemsSt === EntityStatus.loadingSuccess),
      filter(([_getDataBagItemsSt, dataBagItemsState]) =>
        !isNil(dataBagItemsState)),
      takeUntil(this.isDestroyed))
      .subscribe(([_getDataBagItemsSt, dataBagItemsState]) => {
        this.dataBagItems = dataBagItemsState.items;
        this.total = dataBagItemsState.total;
        this.appendActiveToItems(this.dataBagItems);
        this.dataBagsDetailsLoading = false;
        this.searching = false;
      });

    this.store.select(deleteStatus).pipe(
      filter(status => status === EntityStatus.loadingSuccess),
      takeUntil(this.isDestroyed))
      .subscribe(() => {
        this.searching = true;
        if (this.dataBagItems &&
          this.dataBagItems.length === 0 &&
          this.current_page !== 1) {
          this.current_page = this.current_page - 1;
        }
        this.getDataBagItemsData();
      });

    combineLatest([
      this.store.select(getStatus),
      this.store.select(dataBagItemDetailsFromRoute)
    ]).pipe(
      filter(([getDataBagItemDetailsSt, _dataBagItemDetailsState]) =>
        getDataBagItemDetailsSt === EntityStatus.loadingSuccess),
      filter(([_getDataBagItemDetailsSt, dataBagItemDetailsState]) =>
        !isNil(dataBagItemDetailsState)),
      takeUntil(this.isDestroyed))
      .subscribe(([_getDataBagItemDetailsSt, dataBagItemDetailsState]) => {
        this.selectedItemDetails = JSON.parse(dataBagItemDetailsState.data);
        this.dataBagsItemDetailsLoading = false;
      });
  }

  public handleItemSelected(item: string, index: number): void {
    this.selectedItem = item;
    this.dataBagsItemDetailsLoading = true;
    this.store.dispatch(new GetDataBagItemDetails({
      server_id: this.serverId,
      org_id: this.orgId,
      name: this.dataBagName,
      item_name: item
    }));

    this.dataBagItems.filter(
      (d_item, i) => i !== index && d_item.active
    ).forEach(menu => menu.active = !menu.active);

    this.dataBagItems[index].active = !this.dataBagItems[index].active;
    this.activeClassName = 'autoHeight';
  }

  refreshData(data: string) {
    if (data) {
      this.selectedItemDetails = JSON.parse(data);
    }
  }

  ngOnDestroy(): void {
    this.isDestroyed.next(true);
    this.isDestroyed.complete();
  }

  appendActiveToItems (items: DataBagItems[]) {
    items.forEach((i, index) => {
      const tempItem = i;
      tempItem['active'] = false;
      items[index] = tempItem;
    });
    this.dataBagItems = items;
  }

  searchDataBagItems(currentText: string) {
    this.searching = true;
    this.current_page = 1;
    this.searchValue = currentText;
    this.getDataBagItemsData();
  }

  onPageChange(event: number): void {
    this.searching = true;
    this.current_page = event;
    this.getDataBagItemsData();
  }

  getDataBagItemsData() {
    const payload = {
      databagName: this.searchValue,
      server_id: this.serverId,
      org_id: this.orgId,
      name: this.dataBagName,
      page: this.current_page,
      per_page: this.per_page
    };
    this.store.dispatch(new GetDataBagItems(payload));
  }

  public startDataBagItemDelete(dataBagItem: DataBagItems): void {
    this.dataBagItemToDelete = dataBagItem;
    this.deleteModalVisible = true;
  }

  public deleteDataBagItem(): void {
    this.closeDeleteModal();
    this.store.dispatch(new DeleteDataBagItem({
      server_id: this.serverId,
      org_id: this.orgId,
      databag_name: this.dataBagName,
      name: this.dataBagItemToDelete.name
    }));
  }

  public closeDeleteModal(): void {
    this.deleteModalVisible = false;
  }

  public openCreateModal(): void {
    this.openDataBagModal.emit();
  }

  public startUpdateDataBagItem(item: DataBagItems, jsonData: Object): void {
    this.dataBagItemName = item.name;
    this.itemDataJson = JSON.stringify(jsonData, null, 4);
    this.openEditDataBagItemModal.emit();
  }

  openDatabagItemModal(): void {
    this.openDataBagItemModal.emit();
  }
}
