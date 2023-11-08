// import { readDir, BaseDirectory, readTextFile, exists } from '@tauri-apps/api/fs';
import { DragDropContext, Draggable, Droppable } from 'react-beautiful-dnd';
// import { appWindow, currentMonitor } from '@tauri-apps/api/window';
// import { appConfigDir, join } from '@tauri-apps/api/path';
// import { convertFileSrc } from '@tauri-apps/api/tauri';
import { Spacer, Button } from '@nextui-org/react';
import { AiFillCloseCircle } from 'react-icons/ai';
import React, { useState, useEffect } from 'react';
import { BsPinFill } from 'react-icons/bs';
import LanguageArea from './components/LanguageArea';
import SourceArea from './components/SourceArea';
import TargetArea from './components/TargetArea';
import { useConfig } from '../../hooks';
let blurTimeout = null;
let resizeTimeout = null;
let moveTimeout = null;
let osType = "Windows_NT"
// import { store } from '../../utils/store';

export default function Translate() {

    const [alwaysOnTop] = useConfig('translate_always_on_top', false);
    const [hideSource] = useConfig('hide_source', false);
    const [hideLanguage] = useConfig('hide_language', false);
    const [pined, setPined] = useState(false);
    const [serviceConfig, setServiceConfig] = useState(null);
    // const [translateServiceList, setTranslateServiceList] = useState([
    //     'deepl',
    //     'bing',
    //     'yandex',
    //     'google',
    // ]);
    const [translateServiceList, setTranslateServiceList] = useConfig('translate_service_list', [
        // 'deepl',
        'youdao',
        // 'yandex',
        // 'google',
    ]);

    const reorder = (list, startIndex, endIndex) => {
        const result = Array.from(list);
        const [removed] = result.splice(startIndex, 1);
        result.splice(endIndex, 0, removed);
        return result;
    };

    const [pluginList, setPluginList] = useState(null);

    const onDragEnd = async (result) => {
        if (!result.destination) return;
        const items = reorder(translateServiceList, result.source.index, result.destination.index);
        setTranslateServiceList(items);
    };

    const getServiceConfig = async () => {
        let config = {};

        for (const service of translateServiceList) {
            console.log(service)
            config[service] = {}
            // config[service] = (await store.get(service)) ?? {};
        }
        setServiceConfig({ ...config });
    };

    useEffect(() => {
        if (translateServiceList !== null) {
            getServiceConfig();
        }
    }, [translateServiceList]);

    // 是否默认置顶
    useEffect(() => {
        if (alwaysOnTop !== null && alwaysOnTop) {
            // appWindow.setAlwaysOnTop(true);
            // unlistenBlur();
            setPined(true);
        }
    }, [alwaysOnTop]);

    return (
        (
            <div
                className={`bg-background h-screen w-screen ${osType === 'Linux' && 'rounded-[10px] border-1 border-default-100'
                    }`}
            >
                <div
                    className='fixed top-[5px] left-[5px] right-[5px] h-[30px]'
                    data-tauri-drag-region='true'
                />
                <div className={`h-[35px] w-full flex ${osType === 'Darwin' ? 'justify-end' : 'justify-between'}`}>
                    <Button
                        isIconOnly
                        size='sm'
                        variant='flat'
                        disableAnimation
                        className='my-auto bg-transparent'
                        onPress={() => {
                            if (pined) {
                                if (closeOnBlur) {
                                    unlisten = listenBlur();
                                }
                                appWindow.setAlwaysOnTop(false);
                            } else {
                                unlistenBlur();
                                appWindow.setAlwaysOnTop(true);
                            }
                            setPined(!pined);
                        }}
                    >
                        <BsPinFill className={`text-[20px] ${pined ? 'text-primary' : 'text-default-400'}`} />
                    </Button>
                    <Button
                        isIconOnly
                        size='sm'
                        variant='flat'
                        disableAnimation
                        className={`my-auto ${osType === 'Darwin' && 'hidden'} bg-transparent`}
                        onPress={() => {
                            void appWindow.close();
                        }}
                    >
                        <AiFillCloseCircle className='text-[20px] text-default-400' />
                    </Button>
                </div>
                <div className={`${osType === 'Linux' ? 'h-[calc(100vh-37px)]' : 'h-[calc(100vh-35px)]'} px-[8px]`}>
                    <div className='h-full overflow-y-auto'>
                        <div className={`${hideSource && 'hidden'}`}>
                            <SourceArea pluginList={pluginList} />
                            <Spacer y={2} />
                        </div>

                        <div className={`${hideLanguage && 'hidden'}`}>
                            <LanguageArea />
                            <Spacer y={2} />
                        </div>

                        <DragDropContext onDragEnd={onDragEnd}>
                            <Droppable
                                droppableId='droppable'
                                direction='vertical'
                            >
                                {(provided) => (
                                    <div
                                        ref={provided.innerRef}
                                        {...provided.droppableProps}
                                    >
                                        {translateServiceList !== null &&
                                            serviceConfig !== null &&
                                            translateServiceList.map((service, index) => {
                                                const config = serviceConfig[service] ?? {};
                                                const enable = config['enable'] ?? true;
                                                return enable ? (
                                                    <Draggable
                                                        key={service}
                                                        draggableId={service}
                                                        index={index}
                                                    >
                                                        {(provided) => (
                                                            <div
                                                                ref={provided.innerRef}
                                                                {...provided.draggableProps}
                                                            >
                                                                <TargetArea
                                                                    {...provided.dragHandleProps}
                                                                    pluginList={pluginList}
                                                                    name={service}
                                                                    index={index}
                                                                    translateServiceList={translateServiceList}
                                                                />
                                                                <Spacer y={2} />
                                                            </div>
                                                        )}
                                                    </Draggable>
                                                ) : (
                                                    <></>
                                                );
                                            })}
                                    </div>
                                )}
                            </Droppable>
                        </DragDropContext>
                    </div>
                </div>
            </div>
        )
    );
}
