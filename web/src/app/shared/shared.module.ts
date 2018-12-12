import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {ClarityModule, ClrFormsNextModule, ClrModalModule} from '@clr/angular';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

@NgModule({
  imports: [
    CommonModule,
    ClrModalModule,
    ClrFormsNextModule,
    FormsModule,
    ReactiveFormsModule,
    ClarityModule
  ],
  declarations: []
})
export class SharedModule {
}
