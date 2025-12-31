import { Component, OnInit } from '@angular/core';
import { MenuItem, MessageService } from "primeng/api";
import { ToastModule } from "primeng/toast";
import { MenubarModule } from "primeng/menubar";
import { BadgeModule } from "primeng/badge";
import { NgClass } from "@angular/common";
import { AvatarModule } from "primeng/avatar";
import { Router } from "@angular/router";
import { AuthService } from "../services/auth/auth.service";

@Component({
  selector: 'app-menu',
  standalone: true,
  imports: [
    ToastModule,
    MenubarModule,
  ],
  providers: [MessageService],
  templateUrl: './menu.component.html',
  styleUrl: './menu.component.css'
})
export class MenuComponent implements OnInit {
  items: MenuItem[] | undefined;

  constructor(private messageService: MessageService, private router: Router, private authService: AuthService) { }

  ngOnInit() {
    this.authService.isAuthenticated$.subscribe((isAuthenticated) => {
      this.updateMenu(isAuthenticated);
    });
  }

  updateMenu(isAuthenticated: boolean) {
    if (isAuthenticated) {
      this.items = [
        {
          label: 'Dashboard',
          icon: 'pi pi-home',
          command: () => {
            this.router.navigate(['/dashboard']);
          }
        },
        {
          label: 'Add URL',
          icon: 'pi pi-plus',
          command: () => {
            this.router.navigate(['/dashboard']);
          }
        },
        {
          label: 'Logout',
          icon: 'pi pi-sign-out',
          command: () => {
            this.authService.logout().subscribe(() => {
              this.authService.setAuthenticated(false);
              this.router.navigate(['/']);
              this.messageService.add({ severity: 'success', summary: 'Success', detail: 'Logged out' });
            });
          }
        }
      ];
    } else {
      this.items = [
        {
          label: 'Home',
          icon: 'pi pi-home',
          command: () => {
            this.router.navigate(['/']);
          }
        },
        {
          label: 'Login',
          icon: 'pi pi-sign-in',
          command: () => {
            this.router.navigate(['/login']);
          }
        },
        {
          label: 'Sign Up',
          icon: 'pi pi-user-plus',
          command: () => {
            this.router.navigate(['/signup']);
          }
        },
        {
          label: 'About',
          icon: 'pi pi-info-circle',
        }
      ];
    }
  }

}
