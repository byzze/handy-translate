// import { fetch } from '@tauri-apps/api/http';
// import { store } from '../../../utils/store';

export async function tts(text, lang, options = {}) {
    const { config } = options;

    let lingvaConfig = { requestPath: 'lingva.pot-app.com' };

    if (config !== undefined) {
        lingvaConfig = config;
    }

    let { requestPath } = lingvaConfig;
    if (!requestPath.startsWith('http')) {
        requestPath = 'https://' + requestPath;
    }

    const response = await fetch(`${requestPath}/api/v1/audio/${lang}/${encodeURIComponent(text)}`);

    const jsonData = await response.json();

    if (response.ok) {
        return jsonData['audio'];
    }
}

export * from './Config';
export * from './info';
