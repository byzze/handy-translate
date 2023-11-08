import { createRoot } from 'react-dom/client'
import React from 'react'
import ReactDOM from 'react-dom/client';
import { NextUIProvider } from '@nextui-org/react';
import { ThemeProvider as NextThemesProvider } from 'next-themes';
import App from './App'
import './style.css'

// import { initStore } from './utils/store';
// import { initAppVersion, initOsVersion, initOsType, initArch } from './utils/env';

// initStore().then(async () => {
// await initOsType();
// await initArch();
// await initOsVersion();
// await initAppVersion();
// const rootElement = document.getElementById('root');
// const root = ReactDOM.createRoot(rootElement);
// root.render(
//     <NextUIProvider>
//         <NextThemesProvider attribute='class'>
//             <App />
//         </NextThemesProvider>
//     </NextUIProvider>
// );
// });

const container = document.getElementById('root')

const root = createRoot(container)
root.render(
    <NextUIProvider>
        <NextThemesProvider attribute='class'>
            <App />
        </NextThemesProvider>
    </NextUIProvider>
);

