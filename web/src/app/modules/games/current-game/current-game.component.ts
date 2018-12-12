import {Component, OnInit} from '@angular/core';
import {GameService} from '../shared/game.service';
import {Game} from '../../../shared/models/game.model';
import {Cell} from '../../../shared/models/cell.model';

@Component({
  selector: 'app-current-game',
  templateUrl: './current-game.component.html',
  styleUrls: ['./current-game.component.scss']
})
export class CurrentGameComponent implements OnInit {

  currentGame: Game;

  constructor(private gameService: GameService) {
    this.currentGame = this.gameService.currentGame.value;
  }

  ngOnInit() {
  }

  onClickedCell(cell: Cell) {
    this.gameService
      .clickCell(this.currentGame.id, this.currentGame.email, cell)
      .subscribe(game => this.currentGame = game);
  }

  onFlaggedCell(cell: Cell) {
    this.gameService
      .flaggedCell(this.currentGame.id, this.currentGame.email, cell)
      .subscribe(game => this.currentGame = game);
  }
}
