import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SnackAlertComponent } from './snack-alert.component';

describe('SnackAlertComponent', () => {
  let component: SnackAlertComponent;
  let fixture: ComponentFixture<SnackAlertComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SnackAlertComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SnackAlertComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
