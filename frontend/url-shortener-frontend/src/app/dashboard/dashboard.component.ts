import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ButtonModule } from 'primeng/button';
import { InputTextModule } from 'primeng/inputtext';
import { TableModule } from 'primeng/table';
import { CardModule } from 'primeng/card';
import { UrlService } from '../services/url/url.service';
import { MessageService } from 'primeng/api';
import { ToastModule } from 'primeng/toast';
import { TooltipModule } from 'primeng/tooltip';

@Component({
    selector: 'app-dashboard',
    standalone: true,
    imports: [
        CommonModule,
        FormsModule,
        ButtonModule,
        InputTextModule,
        TableModule,
        CardModule,
        ToastModule,
        TooltipModule
    ],
    providers: [MessageService],
    templateUrl: './dashboard.component.html',
    styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
    fullUrl: string = '';
    links: any[] = [];
    loading: boolean = false;

    constructor(private urlService: UrlService, private messageService: MessageService) { }

    ngOnInit(): void {
        this.getAllLinks();
    }

    getAllLinks() {
        this.loading = true;
        this.urlService.getLinks().subscribe({
            next: (data) => {
                this.links = data || [];
                this.loading = false;
            },
            error: (err) => {
                console.error(err);
                this.loading = false;
            }
        });
    }

    addUrl() {
        if (!this.fullUrl) return;

        this.urlService.addUrl(this.fullUrl).subscribe({
            next: (res) => {
                this.messageService.add({ severity: 'success', summary: 'Success', detail: 'URL Shortened!' });
                this.fullUrl = '';
                this.getAllLinks();
            },
            error: (err) => {
                this.messageService.add({ severity: 'error', summary: 'Error', detail: 'Failed to shorten URL' });
                console.error(err);
            }
        });
    }

    deleteUrl(id: number) {
        this.urlService.deleteLink(id).subscribe({
            next: (res) => {
                this.messageService.add({ severity: 'success', summary: 'Success', detail: 'URL Deleted' });
                this.getAllLinks();
            },
            error: (err) => {
                this.messageService.add({ severity: 'error', summary: 'Error', detail: 'Failed to delete URL' });
                console.error(err);
            }
        });
    }

    copyToClipboard(shortUrl: string) {
        // Ideally use domain + shortUrl
        const fullShortUrl = `http://localhost:8080/${shortUrl}`; // Assuming backend handles redirection from root
        navigator.clipboard.writeText(fullShortUrl).then(() => {
            this.messageService.add({ severity: 'info', summary: 'Copied', detail: 'Link copied to clipboard' });
        });
    }
}
