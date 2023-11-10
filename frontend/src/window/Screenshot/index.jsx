import React, { useEffect, useState, useRef } from 'react';
// import { appCacheDir, join } from '@tauri-apps/api/path';
// import { currentMonitor } from '@tauri-apps/api/window';
// import { convertFileSrc } from '@tauri-apps/api/tauri';
// import { appWindow } from '@tauri-apps/api/window';
// import { emit } from '@tauri-apps/api/event';
// import { warn } from 'tauri-plugin-log-api';
// import { invoke } from '@tauri-apps/api';

import { WindowHide, WindowUnfullscreen, WindowFullscreen, WindowMaximise, WindowSetBackgroundColour, WindowShow, WindowSetSize, WindowSetAlwaysOnTop, WindowUnmaximise, WindowMinimise, EventsOn, EventsEmit, ClipboardSetText, Hide } from "../../../wailsjs/runtime"

export default function Screenshot() {
    //screenshot/screenshot-1699605146407919600.png
    const [imgurl, setImgurl] = useState('screenshot/screenshot.png');
    const [isMoved, setIsMoved] = useState(false);
    const [isDown, setIsDown] = useState(false);
    const [mouseDownX, setMouseDownX] = useState(0);
    const [mouseDownY, setMouseDownY] = useState(0);
    const [mouseMoveX, setMouseMoveX] = useState(0);
    const [mouseMoveY, setMouseMoveY] = useState(0);

    const imgRef = useRef();

    const canvasRef = useRef(null);

    const captureScreenshot = (x, y, width, height) => {
        const canvas = canvasRef.current;
        const context = canvas.getContext('2d');

        // 创建一个新的 Image 对象
        const image = new Image();
        image.src = 'screenshot/screenshot.png'; // 替换为您的图片 URL

        image.onload = function () {
            // 设置截图的起始坐标和截图的宽度和高度
            // const x = 100; // 起始横坐标
            // const y = 50; // 起始纵坐标
            // const width = 200; // 截图宽度
            // const height = 150; // 截图高度

            // 在 Canvas 上绘制截图
            canvas.width = width;
            canvas.height = height;
            context.drawImage(image, x, y, width, height, 0, 0, width, height);
            const base64Data = canvas.toDataURL('image/png');
            console.log(base64Data)
        };
    }

    useEffect(() => {
        // EventsOn("screenshot", (result) => {
        //     WindowMaximise()
        //     setImgurl(result)
        // })
        WindowMaximise()
        // WindowUnmaximise()
        console.log(mouseDownX, mouseDownY, mouseMoveX, mouseMoveY)
        // const position = monitor.position;
        // console.log(monitor.position)
        // currentMonitor().then((monitor) => {
        //     const position = monitor.position;
        //     invoke('screenshot', { x: position.x, y: position.y }).then(() => {
        //         appCacheDir().then((appCacheDirPath) => {
        //             join(appCacheDirPath, 'pot_screenshot.png').then((filePath) => {
        //                 setImgurl(convertFileSrc(filePath));
        //             });
        //         });
        //     });
        // });
    }, []);

    const keyDown = (event) => {
        if (event.key === 'Escape') {
            WindowUnmaximise();
        }
    };

    return (
        <>
            <canvas ref={canvasRef}></canvas>
            <img
                ref={imgRef}
                className='fixed top-0 left-0 w-full select-none'
                src={imgurl}
                draggable={false}
                onLoad={() => {
                    if (imgurl !== '' && imgRef.current.complete) {
                        // void WindowShow();
                        // void appWindow.setFocus();
                        // void appWindow.setResizable(false);
                    }
                }}
            />
            <div
                className={`fixed bg-[#2080f020] border border-solid border-sky-500 ${!isMoved && 'hidden'}`}
                style={{
                    top: Math.min(mouseDownY, mouseMoveY),
                    left: Math.min(mouseDownX, mouseMoveX),
                    bottom: screen.height - Math.max(mouseDownY, mouseMoveY),
                    right: screen.width - Math.max(mouseDownX, mouseMoveX),
                }}
            />
            <div
                className='fixed top-0 left-0 bottom-0 right-0 cursor-crosshair select-none'
                onMouseDown={(e) => {
                    if (e.buttons === 1) {
                        setIsDown(true);
                        setMouseDownX(e.clientX);
                        setMouseDownY(e.clientY);
                    } else {
                        // void appWindow.close();
                    }
                }}
                onMouseMove={(e) => {
                    if (isDown) {
                        setIsMoved(true);
                        setMouseMoveX(e.clientX);
                        setMouseMoveY(e.clientY);
                    }
                }}
                onMouseUp={async (e) => {
                    // WindowHide();
                    setIsDown(false);
                    setIsMoved(false);

                    console.log({ mouseDownX })
                    console.log(screen.width)
                    const imgWidth = imgRef.current.naturalWidth;
                    const dpi = imgWidth / screen.width;
                    const left = Math.floor(Math.min(mouseDownX, e.clientX) * dpi);
                    const top = Math.floor(Math.min(mouseDownY, e.clientY) * dpi);
                    const right = Math.floor(Math.max(mouseDownX, e.clientX) * dpi);
                    const bottom = Math.floor(Math.max(mouseDownY, e.clientY) * dpi);
                    const width = right - left;
                    const height = bottom - top;
                    captureScreenshot(left, top, width, height)
                    console.log({ left, top, width, height })
                    if (width <= 0 || height <= 0) {
                        // warn('Screenshot area is too small');
                        // await appWindow.close();
                    } else {
                        // await invoke('cut_image', { left, top, width, height });
                        // await emit('success');
                        // await appWindow.close();
                    }
                }}
            />
        </>
    );
}
