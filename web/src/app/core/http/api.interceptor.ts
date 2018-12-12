import {Injectable} from '@angular/core';
import {HttpEvent, HttpInterceptor, HttpHandler, HttpRequest} from '@angular/common/http';
import {Observable, throwError} from 'rxjs';
import {catchError, map} from 'rxjs/operators';
import {environment} from '../../../environments/environment';

@Injectable()
export class APIInterceptor implements HttpInterceptor {

  constructor() {
  }

  static updateUrl(req: string): string {
    return environment.apiUrl + req;
  }

  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    request = request.clone({
      url: APIInterceptor.updateUrl(request.url)
    });

    return next.handle(request)
      .pipe(
        map((response) => {
          return response;
        }),
        catchError((response) => {
          return throwError(response);
        })
      );
  }

}
