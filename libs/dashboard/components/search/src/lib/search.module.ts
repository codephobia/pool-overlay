import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

import { SearchComponent } from './search.component';

@NgModule({
    imports: [CommonModule, FormsModule, ReactiveFormsModule, FontAwesomeModule],
    declarations: [SearchComponent],
    exports: [SearchComponent],
})
export class SearchModule { }
