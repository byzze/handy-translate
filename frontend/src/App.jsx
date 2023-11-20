import { useState, useEffect } from 'react';
import logo from './assets/images/logo-universal.png';
import { BrowserRouter } from 'react-router-dom';
import { useTranslation } from 'react-i18next';
import { useTheme } from 'next-themes';

import Translate from './window/Translate';
import Screenshot from './window/Screenshot';
import Home from './window/Home';
import { useConfig, useSyncAtom } from './hooks';
import { EventsOn, WindowGetSize, WindowSetSize, WindowFullscreen, WindowUnfullscreen, WindowIsFullscreen } from '../wailsjs/runtime';
import './i18n';
import './style.css';



function App() {
    const [appTheme] = useConfig('app_theme', 'system');
    const [appLanguage] = useConfig('app_language', 'zh_cn');
    const { setTheme } = useTheme();
    const { i18n } = useTranslation();
    const [appLabel, setAppLabel] = useState('translate');
    const { width, setWidth } = useState(0)
    const { height, setHeight } = useState(0)

    // 共享的变量状态和更新方法
    const [sharedVariable, setSharedVariable] = useState('');

    // 更新共享变量的方法
    const updateSharedVariable = (newValue) => {
        setSharedVariable(newValue);
        // setAppLabel(newValue)
    };

    const windowMap = {
        translate: <Translate variable={sharedVariable} onUpdateVariable={updateSharedVariable} />,
        screenshot: <Screenshot />,
        home: <Home />,
    };

    useEffect(() => {
        console.log(sharedVariable)
    }, [sharedVariable])


    useEffect(() => {
        EventsOn("appLabel", (result) => {
            setAppLabel(result)
        })
    }, [])

    useEffect(() => {
        if (width == 0 && height == 0) {
            if (appLabel == 'translate') {
                WindowGetSize().then((w, h) => {
                    setWidth(w)
                    setHeight(h)
                })
            }
        }
    }, [])

    useEffect(() => {
        if (appLabel == 'translate') {
            WindowIsFullscreen().then((r) => {
                if (r) {
                    WindowUnfullscreen()
                }
            })
            WindowSetSize(460, 460)
        }

        if (appLabel == 'screenshot') {
            WindowFullscreen()
        }

        if (appLabel == 'home') {
            WindowSetSize(666, 666)
        }

    }, [appLabel]);

    useEffect(() => {
        if (appTheme !== null) {
            if (appTheme !== 'system') {
                setTheme(appTheme);
            } else {
                try {
                    if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
                        setTheme('dark');
                    } else {
                        setTheme('light');
                    }
                    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
                        if (e.matches) {
                            setTheme('dark');
                        } else {
                            setTheme('light');
                        }
                    });
                } catch {
                    warn("Can't detect system theme.");
                }
            }
        }
        if (appLanguage !== null) {
            i18n.changeLanguage(appLanguage);
        }
    }, [appTheme, appLanguage]);

    return <BrowserRouter>
        {windowMap[appLabel]}
    </BrowserRouter>;
}

export default App
