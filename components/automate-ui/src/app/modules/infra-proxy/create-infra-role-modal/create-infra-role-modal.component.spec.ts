import { ReactiveFormsModule, FormBuilder } from '@angular/forms';
import { StoreModule, Store } from '@ngrx/store';
import { NgrxStateAtom, ngrxReducers, runtimeChecks } from 'app/ngrx.reducers';
import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';
import { MockComponent } from 'ng2-mock-component';
import { CreateInfraRoleModalComponent } from './create-infra-role-modal.component';
import { TelemetryService } from 'app/services/telemetry/telemetry.service';
import { CreateRoleSuccess, CreateRoleFailure, CreateRolePayload } from 'app/entities/infra-roles/infra-role.action';
import { HttpErrorResponse } from '@angular/common/http';
import { HttpStatus } from 'app/types/types';
import { EventEmitter } from '@angular/core';
import { HttpClient, HttpHandler } from '@angular/common/http';

class MockTelemetryService {
  track() { }
}

describe('CreateInfraRoleModalComponent', () => {
  let component: CreateInfraRoleModalComponent;
  let fixture: ComponentFixture<CreateInfraRoleModalComponent>;

  beforeEach( waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [
        MockComponent({ selector: 'chef-button', inputs: ['disabled'] }),
        MockComponent({ selector: 'chef-loading-spinner' }),
        MockComponent({ selector: 'chef-form-field' }),
        MockComponent({ selector: 'chef-error' }),
        MockComponent({ selector: 'chef-toolbar' }),
        MockComponent({ selector: 'app-infra-tab-change' }),
        MockComponent({ selector: 'app-infra-tab', inputs: ['active', 'disabled'] }),
        MockComponent({ selector: 'chef-modal',
          inputs: ['visible']
        }),
        CreateInfraRoleModalComponent
      ],
      providers: [
        { provide: TelemetryService, useClass: MockTelemetryService },
        HttpClient, HttpHandler
      ],
      imports: [
        ReactiveFormsModule,
        StoreModule.forRoot(ngrxReducers, { runtimeChecks })
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateInfraRoleModalComponent);
    component = fixture.componentInstance;
    component.detailsFormGroup = new FormBuilder().group({
      name: ['', null],
      description: ['', null]
    });
    component.defaultAttrFormGroup = new FormBuilder().group({
      default: ['', null]
    });
    component.overrideAttrFormGroup = new FormBuilder().group({
      override: ['', null]
    });
    component.openEvent = new EventEmitter();
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  describe('create role', () => {
    let store: Store<NgrxStateAtom>;
    const role: CreateRolePayload  = {
      org_id: 'chef_manage',
      server_id: 'test',
      name: 'chef-test-role',
      description: 'test role',
      run_list: ['role[chef-load-role-580613600]', 'role[chef-load-role-37386300]'],
      default_attributes: '{test:test}',
      override_attributes: '{test:test}'
    };

    beforeEach(() => {
        store = TestBed.inject(Store);
    });

    it('opening create modal resets default tab to details tab', () => {
      component.visible = true;
      expect(component.detailsTab).toBe(true);
      expect(component.constraintsTab).toBe(false);
      expect(component.defaultTab).toBe(false);
      expect(component.overrideTab).toBe(false);
    });

    it('opening create modal resets name, description, default, override to empty string', () => {
      component.visible = true;
      expect(component.detailsFormGroup.controls['name'].value).toEqual('');
      expect(component.detailsFormGroup.controls['description'].value).toEqual('');
      expect(component.defaultAttrFormGroup.controls['default'].value).toEqual('');
      expect(component.overrideAttrFormGroup.controls['override'].value).toEqual('');
    });

    it('on conflict error, modal remains open and displays conflict error', () => {
      spyOn(component.conflictErrorEvent, 'emit');
      component.visible = true;
      component.detailsFormGroup.controls['name'].setValue(role.name);
      component.detailsFormGroup.controls['description'].setValue(role.description);
      component.selectedRunList = role.run_list;
      component.defaultAttrFormGroup.controls['default'].setValue(
        JSON.stringify(role.default_attributes));
      component.overrideAttrFormGroup.controls['override'].setValue(
        JSON.stringify(role.override_attributes));
      component.createRole();
      const conflict = <HttpErrorResponse>{
        status: HttpStatus.CONFLICT,
        ok: false
      };
      store.dispatch(new CreateRoleFailure(conflict));
      expect(component.conflictError).toBe(true);
    });

    it('on success, closes modal and adds new role', () => {
      spyOn(component.conflictErrorEvent, 'emit');
      component.visible = true;
      component.detailsFormGroup.controls['name'].setValue(role.name);
      component.detailsFormGroup.controls['description'].setValue(role.description);
      component.selectedRunList = role.run_list;
      component.defaultAttrFormGroup.controls['default'].setValue(
        JSON.stringify(role.default_attributes));
      component.overrideAttrFormGroup.controls['override'].setValue(
        JSON.stringify(role.override_attributes));
      component.createRole();
      store.dispatch(new CreateRoleSuccess(role));
      expect(component.creating).toBe(false);
      expect(component.visible).toBe(false);
    });

    it('on create error, modal is closed (because error is handled by failure banner)', () => {
      spyOn(component.conflictErrorEvent, 'emit');
      component.detailsFormGroup.controls['name'].setValue(role.name);
      component.detailsFormGroup.controls['description'].setValue(role.description);
      component.selectedRunList = role.run_list;
      component.defaultAttrFormGroup.controls['default'].setValue(
        JSON.stringify(role.default_attributes));
      component.overrideAttrFormGroup.controls['override'].setValue(
        JSON.stringify(role.override_attributes));
      component.createRole();
      const error = <HttpErrorResponse>{
        status: HttpStatus.INTERNAL_SERVER_ERROR,
        ok: false
      };
      store.dispatch(new CreateRoleFailure(error));
      expect(component.conflictError).toBe(false);
    });
  });
});
