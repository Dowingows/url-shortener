import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    vus: 150,           // usuários virtuais simultâneos
    duration: '30s',   // duração do teste
};

export default function () {
    const url = 'http://localhost:9999/shortener';

    const payload = JSON.stringify({
        url: 'https://google.com'
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    const res = http.post(url, payload, params);

    check(res, {
        'status 201 ou 200': (r) => r.status === 201 || r.status === 200,
        'retorna short_url': (r) => r.body.includes('short_url'),
    });

    sleep(0.5); // espera 1 segundo entre requisições
}
