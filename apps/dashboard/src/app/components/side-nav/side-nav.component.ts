import { Component } from '@angular/core';

@Component({
    selector: 'pool-overlay-side-nav',
    templateUrl: './side-nav.component.html',
})
export class SideNavComponent {
    public buttons = [{
        title: 'Players',
        link: '#',
    }, {
        title: 'Teams',
        link: '#',
    }, {
        title: 'Games',
        link: '#',
    }];
}
