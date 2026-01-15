import http from 'k6/http';
import { check, sleep } from 'k6';

// Configuração de cenários
export let options = {
    scenarios: {
        create_urls: {
            executor: 'per-vu-iterations',
            vus: 1500,
            iterations: 10,
            maxDuration: '30s',
            exec: 'createURL', // função que será chamada
        },
        consume_urls: {
            executor: 'constant-vus',
            vus: 15000,
            duration: '60s',
            startTime: '30s',
            exec: 'consumeURL', // função que será chamada
        },
    },
};

// Array global para armazenar short URLs
let shortUrls = [];

// ✅ Função de criação de URL
export function createURL() {
    const url = 'http://localhost:9999/shortener';
    const payload = JSON.stringify({ url: 'https://google.com' });
    const params = { headers: { 'Content-Type': 'application/json' } };

    const res = http.post(url, payload, params);

    check(res, {
        'POST status 200 ou 201': (r) => r.status === 200 || r.status === 201,
        'retorna short_url': (r) => r.body.includes('short_url'),
    });

    // Armazena a short URL globalmente
    const shortUrl = JSON.parse(res.body).short_url;
    shortUrls.push(shortUrl);
    console.log(`VU ${__VU} criou short URL: ${shortUrl}`);

    sleep(1);
}

// ✅ Função de consumo da URL
export function consumeURL() {
    // Cada VU consome todas as URLs criadas até agora
    shortUrls.forEach((url) => {
        const res = http.get(url);

        check(res, {
            'GET status 200 ou 302': (r) => r.status === 200 || r.status === 302,
        });
    });

    sleep(0.5);
}
