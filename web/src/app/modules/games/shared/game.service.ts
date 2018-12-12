import {Injectable} from '@angular/core';
import {BehaviorSubject, Observable} from 'rxjs';
import {Game} from '../../../shared/models/game.model';
import {HttpClient} from '@angular/common/http';
import {map} from 'rxjs/operators';
import {Cell} from '../../../shared/models/cell.model';

@Injectable({
  providedIn: 'root'
})
export class GameService {

  private url = '/games';
  currentGame: BehaviorSubject<Game>;

  constructor(private http: HttpClient) {
    this.currentGame = new BehaviorSubject(null);
  }

  createGame(email: string, size: number): Observable<Game> {
    return this.http.post<Game>(`${this.url}`, {
      email: email,
      size: size
    }).pipe(map(result => result));
  }

  clickCell(gameId: number, email: string, cell: Cell): Observable<Game> {
    return this.http.patch<Game>(`${this.url}/open`, {
      gameId: gameId,
      email: email,
      cell: cell
    }).pipe(map(result => result));
  }

  flaggedCell(gameId: number, email: string, cell: Cell): Observable<Game> {
    return this.http.patch<Game>(`${this.url}/flag`, {
      gameId: gameId,
      email: email,
      cell: cell
    }).pipe(map(result => result));
  }
}
