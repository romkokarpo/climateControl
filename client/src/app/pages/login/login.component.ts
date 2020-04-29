import { AuthenticationService } from './../../services/authentication/authentication.service';
import { Component, OnInit, OnDestroy } from '@angular/core';
import { Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators, NgForm } from '@angular/forms';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit, OnDestroy {
  userEmail: string;
  userPassword: string;
  rememberUser: boolean;
  submitted = false;

  loginForm;

  constructor(private authService: AuthenticationService, private formBuilder: FormBuilder) {}

  ngOnInit() {
  }
  ngOnDestroy() {
  }

  onSubmit(f: NgForm) {
    this.authService.login(this.userEmail, this.userPassword);
  }
}
