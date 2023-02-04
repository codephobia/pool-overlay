import { Component, EventEmitter, Input, Output } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { faMagnifyingGlass } from '@fortawesome/pro-regular-svg-icons';

let idCounter = 0;

@Component({
    selector: 'dashboard-search',
    templateUrl: './search.component.html'
})
export class SearchComponent {
    @Input()
    public set search(search: string) {
        this._search = search;

        this.form.controls['search'].patchValue(search);
    }
    public get search(): string {
        return this._search;
    }

    @Output()
    public onSearch = new EventEmitter<{ search: string }>();

    public faMagnifyingGlass = faMagnifyingGlass;
    public form: FormGroup;
    public id: string;
    private _search = '';

    constructor(private _fb: FormBuilder) {
        this.id = `dashboard-search-${++idCounter}`;

        this.form = this._fb.group({
            search: '',
        });
    }

    public submit(): void {
        this.onSearch.emit({
            search: '',
        });
    }
}
