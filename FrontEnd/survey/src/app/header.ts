import { HttpClient } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';

export class Header {
    public headers: HttpHeaders = new HttpHeaders();
    constructor() {
        this.headers = this.headers
            .append('username', sessionStorage.getItem('username'))
            .append('password', sessionStorage.getItem('password'));
    }
}
