import { Component } from '@angular/core';
import { faDesktop, faUser, faPool8Ball, faPencil, faTrophyStar } from '@fortawesome/pro-regular-svg-icons';

@Component({
    selector: 'pool-overlay-side-nav',
    templateUrl: './side-nav.component.html',
})
export class SideNavComponent {
    public buttons = [{
        icon: faDesktop,
        title: 'Overlay',
        link: 'overlay',
    }, {
        icon: faTrophyStar,
        title: 'Tournaments',
        link: 'tournaments',
    }, {
        icon: faUser,
        title: 'Players',
        link: 'players',
    }, {
        icon: faPool8Ball,
        title: 'Games',
        link: 'games',
    }, {
        icon: faPencil,
        title: 'Drawing',
        link: 'drawing',
    }];
}
