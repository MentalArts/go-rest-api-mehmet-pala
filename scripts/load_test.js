import http from 'k6/http';
import { check } from 'k6';

export default function () {
  const response = http.get('http://localhost:8080/api/v1/books');
  check(response, {
    'is status 200': (r) => r.status === 200,
  });
}
