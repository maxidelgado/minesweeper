import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {GamesComponent} from './games.component';
import {NewGameComponent} from './new-game/new-game.component';
import {CurrentGameComponent} from './current-game/current-game.component';

const routes: Routes = [
  {path: '', redirectTo: 'new', component: GamesComponent},
  {path: 'new', component: NewGameComponent},
  {path: 'current', component: CurrentGameComponent}
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class GamesRoutingModule {
}
