import { AuthenticationService } from './../../services/authentication/authentication.service';
import { Component, OnInit, OnDestroy } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit, OnDestroy {
  userEmail: string;
  userPassword: string;
  rememberUser: boolean;

  constructor(private authService: AuthenticationService) {}

  ngOnInit() {
  }
  ngOnDestroy() {
  }

  login() {
    this.authService.login(this.userEmail, this.userPassword);
  }
}
