import {Component, OnInit} from '@angular/core';
import {AbstractControl, FormBuilder, FormControl, FormGroup, Validators} from '@angular/forms';
import {GameService} from '../shared/game.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-new-game',
  templateUrl: './new-game.component.html',
  styleUrls: ['./new-game.component.scss']
})
export class NewGameComponent implements OnInit {

  form: FormGroup;

  constructor(private fb: FormBuilder,
              private gameService: GameService,
              private router: Router) {
  }

  ngOnInit() {
    this.createForm();
  }

  createForm() {
    this.form = this.fb.group({
      email: [null, Validators.required],
      cols: [null, Validators.required],
    });
  }

  get email(): AbstractControl {
    return this.form.get('email') as FormControl;
  }

  get cols(): AbstractControl {
    return this.form.get('cols') as FormControl;
  }

  onSubmit() {
    this.gameService
      .createGame(this.email.value, +this.cols.value)
      .subscribe(game => {
        this.gameService.currentGame.next(game);
        this.router.navigateByUrl('/games/current');
      });
  }
}
