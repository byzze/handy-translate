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
    console.log(`${requestPath}/api/v1/audio/${lang}/${encodeURIComponent(text)}`)

    const response = await fetch(`${requestPath}/api/v1/audio/${lang}/${encodeURIComponent(text)}`);
    console.log(response)
    const jsonData = await response.json();
    console.log(jsonData)
    console.log(jsonData['audio'])
    // const res = await fetch(`${requestPath}/api/v1/audio/${lang}/${encodeURIComponent(text)}`);
    // console.log(res)
    if (response.ok) {
        return jsonData['audio'];
    }
}

export * from './Config';
export * from './info';
