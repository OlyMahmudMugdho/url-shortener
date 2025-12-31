import { Component } from '@angular/core';
import { IconFieldModule } from "primeng/iconfield";
import { InputIconModule } from "primeng/inputicon";
import { InputTextModule } from "primeng/inputtext";
import { NgStyle } from "@angular/common";
import { FloatLabelModule } from "primeng/floatlabel";
import { Button, ButtonDirective } from "primeng/button";
import { Ripple } from "primeng/ripple";
import { FormsModule } from "@angular/forms";
import { Router, RouterLink } from "@angular/router";
import { HttpClient } from "@angular/common/http";
import { AuthService } from "../services/auth/auth.service";

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    IconFieldModule,
    InputIconModule,
    InputTextModule,
    NgStyle,
    FloatLabelModule,
    Button,
    ButtonDirective,
    Ripple,
    FormsModule,
    RouterLink
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {
  username: any;
  password: any;

  constructor(private authService: AuthService, private router: Router) {
  }

  handleSubmit() {
    this.authService.sendLoginData({
      username: this.username,
      password: this.password
    }).subscribe({
      next: (response) => {
        if (!response.ok) {
          this.authService.setAuthenticated(false);
          alert("Login failed");
        } else {
          this.authService.setAuthenticated(true);
          this.router.navigate(['/']); // Navigate to home
        }
      },
      error: (err) => {
        this.authService.setAuthenticated(false);
        alert("Login error: " + (err.error?.message || "Unknown error"));
      }
    })
  }
}
