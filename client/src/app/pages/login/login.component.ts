import { AuthenticationService } from './../../services/authentication/authentication.service';
import { Component, OnInit, OnDestroy } from '@angular/core';
import { Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit, OnDestroy {
  userEmail: string;
  userPassword: string;
  rememberUser: boolean;
  loginForm: FormGroup;
  submitted = false;

  constructor(private authService: AuthenticationService, private formBuilder: FormBuilder) {}

  ngOnInit() {
    this.loginForm = this.formBuilder.group({
      userName: ['', Validators.required]
    });
  }
  ngOnDestroy() {
  }

  login() {
    this.authService.login(this.userEmail, this.userPassword);
  }
}
