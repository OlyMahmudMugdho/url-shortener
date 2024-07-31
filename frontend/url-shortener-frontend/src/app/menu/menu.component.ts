import {Component, OnInit} from '@angular/core';
import {MenuItem, MessageService} from "primeng/api";
import {ToastModule} from "primeng/toast";
import {MenubarModule} from "primeng/menubar";
import {BadgeModule} from "primeng/badge";
import {NgClass} from "@angular/common";
import {AvatarModule} from "primeng/avatar";
import {Router} from "@angular/router";

@Component({
  selector: 'app-menu',
  standalone: true,
  imports: [
    ToastModule,
    MenubarModule,
  ],
  providers : [MessageService],
  templateUrl: './menu.component.html',
  styleUrl: './menu.component.css'
})
export class MenuComponent implements OnInit {
  items: MenuItem[] | undefined;
  router: Router

  constructor(private messageService: MessageService, router : Router) {
    this.router = router
  }

  ngOnInit() {
    this.items = [
      {
        command: () => {
          this.router.navigate(['home']);
        },
        icon: 'pi pi-info',
        label: 'About'
      },
      {
        label: 'Login',
        icon: 'pi pi-sign-in',
        command: () => {
          this.messageService.add({ severity: 'warn', summary: 'Search Results', detail: 'No results found', life: 3000 });
        }
      },
      {
        separator: true
      },
      {
        label: 'Add New URL',
        icon: 'pi pi-spin pi-bolt',
      }
    ];
  }

}
