import { Observable, Subscription, timer } from 'rxjs';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss']
})
export class RegisterComponent implements OnInit {
  // private timer: Observable<any>;
  // private subscription: Subscription;

  constructor() { }

  ngOnInit() {
  }

  // public setSuccess() {
  //   this.success = true;

  //   this.timer = timer(5000);
  //   this.subscription = this.timer.subscribe(() => {
  //     this.success = false;
  //   });
  // }
}
