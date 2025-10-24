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

    console.log('TTS API 请求:', `${requestPath}/api/v1/audio/${lang}/${encodeURIComponent(text)}`);
    const response = await fetch(`${requestPath}/api/v1/audio/${lang}/${encodeURIComponent(text)}`);

    const jsonData = await response.json();
    console.log('TTS API 响应类型:', typeof jsonData);
    console.log('TTS API 响应:', jsonData);
    console.log('audio 字段类型:', typeof jsonData?.audio);
    console.log('audio 字段:', jsonData?.audio);

    if (response.ok) {
        let audioData = jsonData['audio'];
        console.log('audioData 类型:', typeof audioData);
        console.log('audioData 长度:', audioData?.length);
        console.log('audioData 是数组吗:', Array.isArray(audioData));
        console.log('audioData constructor:', audioData?.constructor?.name);

        if (audioData) {
            // 如果 audioData 是字符串类型
            if (typeof audioData === 'string') {
                // 移除可能的 data URI 前缀（如 "data:audio/mp3;base64,"）
                if (audioData.includes(',')) {
                    const parts = audioData.split(',');
                    audioData = parts[parts.length - 1];
                    console.log('移除 data URI 前缀后的数据长度:', audioData.length);
                }
            }
            // 如果是对象或数组，直接返回（可能已经是解码后的数据）
            else if (typeof audioData === 'object') {
                console.log('audioData 是对象类型，直接返回');
            }

            return audioData;
        }
    }

    throw new Error(`TTS API 失败: ${response.status} ${response.statusText}`);
}

export * from './info';
