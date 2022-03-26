import { Component } from '@angular/core';

@Component({
    selector: 'pool-overlay-side-nav',
    templateUrl: './side-nav.component.html',
})
export class SideNavComponent {
    public buttons = [{
        title: 'Overlay',
        link: 'overlay',
    }, {
        title: 'Players',
        link: 'players',
    }, {
        title: 'Games',
        link: 'games',
    }, {
        title: 'Drawing',
        link: 'drawing',
    }];
}
