import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { OutletComponent } from '../shared/components/outlet/outlet.component';
import { HomePageComponent } from './containers/home/home-page.component';

const routes: Routes = [
    {
        path: '',
        component: OutletComponent,
        children: [
            {
                path: '',
                pathMatch: 'full',
                component: HomePageComponent,
            },
        ],
    }
];

@NgModule({
    imports: [
        CommonModule,
        RouterModule.forChild(routes),
    ],
    exports: [RouterModule],
})
export class OverlayRoutingModule { }
