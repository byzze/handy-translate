import { useCallback, useEffect } from 'react';

let audioContext = null;
let source = null;

// 初始化 AudioContext
const initAudioContext = () => {
    if (!audioContext) {
        audioContext = new (window.AudioContext || window.webkitAudioContext)();
        console.log('AudioContext 初始化完成，状态:', audioContext.state);
    }

    // 某些浏览器需要用户交互后才能启动 AudioContext
    if (audioContext.state === 'suspended') {
        console.log('AudioContext 被暂停，尝试恢复...');
        audioContext.resume().then(() => {
            console.log('AudioContext 已恢复，状态:', audioContext.state);
        });
    }

    return audioContext;
};

export const useVoice = () => {
    useEffect(() => {
        // 组件挂载时初始化
        initAudioContext();
    }, []);

    const playOrStop = useCallback(async (data) => {
        try {
            const ctx = initAudioContext();

            if (source) {
                // 如果正在播放，停止播放
                console.log('停止当前播放');
                source.stop();
                source.disconnect();
                source = null;
                return;
            }

            // 如果没在播放，开始播放
            console.log('useVoice: 开始解码音频数据，字节长度:', data.length);
            console.log('AudioContext 状态:', ctx.state);

            // 确保 AudioContext 是运行状态
            if (ctx.state === 'suspended') {
                await ctx.resume();
                console.log('AudioContext 已恢复');
            }

            // 使用 Promise 版本的 decodeAudioData
            const buffer = await ctx.decodeAudioData(data.buffer);
            console.log('音频解码成功，时长:', buffer.duration, '秒', '采样率:', buffer.sampleRate);

            source = ctx.createBufferSource();
            source.buffer = buffer;
            source.connect(ctx.destination);

            console.log('开始播放音频...');
            source.start(0);

            source.onended = () => {
                console.log('音频播放完成');
                if (source) {
                    source.disconnect();
                    source = null;
                }
            };

            // 返回 Promise 等待播放完成
            return new Promise((resolve) => {
                const onEnd = () => {
                    console.log('播放结束回调');
                    resolve();
                };
                if (source) {
                    source.addEventListener('ended', onEnd, { once: true });
                } else {
                    resolve();
                }
            });
        } catch (err) {
            console.error('音频播放错误:', err);
            console.error('错误详情:', err.message, err.stack);
            if (source) {
                try {
                    source.disconnect();
                } catch (e) {
                    // 忽略断开连接错误
                }
                source = null;
            }
            throw err; // 抛出错误让调用者处理
        }
    }, []);

    return playOrStop;
};
