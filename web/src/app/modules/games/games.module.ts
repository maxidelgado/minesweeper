import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {GamesComponent} from './games.component';
import {GamesRoutingModule} from './games-routing.module';
import {NewGameComponent} from './new-game/new-game.component';
import {SharedModule} from '../../shared/shared.module';
import {GameService} from './shared/game.service';
import { CurrentGameComponent } from './current-game/current-game.component';
import { CellComponent } from './current-game/cell/cell.component';

@NgModule({
  declarations: [GamesComponent, NewGameComponent, CurrentGameComponent, CellComponent],
  imports: [
    CommonModule,
    GamesRoutingModule,
    SharedModule
  ],
  providers: [GameService]
})
export class GamesModule {
}
