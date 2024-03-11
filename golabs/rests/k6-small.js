import http from 'k6/http';

const url = 'http://localhost:8080/small';

export default function () {
  let data = { name: 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa' };
  http.post(url, JSON.stringify(data), {
    headers: { 'Content-Type': 'application/json' },
  });
}