import React, { useEffect, useRef, useState } from "react";
import { Button, Card, CardBody, CardFooter, ButtonGroup, Chip, Tooltip, Spinner } from '@nextui-org/react';
import toast, { Toaster } from 'react-hot-toast';
import { atom, useAtom } from 'jotai';
import { useTranslation } from 'react-i18next';
import { HiOutlineVolumeUp } from 'react-icons/hi';
import { HiTranslate } from 'react-icons/hi';
import { LuDelete } from 'react-icons/lu';
import { MdContentCopy } from 'react-icons/md';
import { MdSmartButton } from 'react-icons/md';
import { WindowHide, WindowShow, EventsOn, ClipboardSetText } from "../../../../../wailsjs/runtime"

import { useConfig, useSyncAtom, useVoice, useToastStyle } from '../../../../hooks';
import * as builtinTtsServices from '../../../../services/tts';
import detect from '../../../../utils/lang_detect';

import { sourceLanguageAtom } from "../LanguageArea";
import { createWorker } from 'tesseract.js';

export const sourceTextAtom = atom('');
export const detectLanguageAtom = atom('');

export default function SourceArea(props) {
    const { pluginList } = props;
    const [sourceText, setSourceText, syncSourceText] = useSyncAtom(sourceTextAtom);
    const [detectLanguage, setDetectLanguage] = useAtom(detectLanguageAtom);
    const [recognizeLanguage] = useConfig('recognize_language', 'auto');
    const [ttsServiceList] = useConfig('tts_service_list', ['lingva_tts']);
    const [deleteNewline] = useConfig('translate_delete_newline', false);
    const [recognizeServiceList] = useConfig('recognize_service_list', ['system', 'tesseract']);
    const [dynamicTranslate] = useConfig('dynamic_translate', false);
    const toastStyle = useToastStyle();
    const { t } = useTranslation();

    const [isSpeakLoading, setIsSpeakLoading] = useState(false)
    const textAreaRef = useRef();
    const speak = useVoice();

    useEffect(() => {
        textAreaRef.current.style.height = '50px';
        textAreaRef.current.style.height = textAreaRef.current.scrollHeight + 'px';
    }, [sourceText]);


    const detect_language = async (text) => {
        setDetectLanguage(await detect(text));
    };

    const keyDown = (event) => {
        if (event.key === 'Enter' && !event.shiftKey) {
            event.preventDefault();
            detect_language(sourceText).then(() => {
                syncSourceText();
            });
        }
        if (event.key === 'Escape') {
            WindowHide();
        }
    };

    useEffect(() => {
        EventsOn("query", (result) => {
            setSourceText(result)
        })
    }, []);

    const handleSpeak = async () => {
        try {
            setIsSpeakLoading(true)

            let lang = await detect(sourceText);
            setDetectLanguage(lang)

            const serviceName = ttsServiceList[0];
            if (serviceName.startsWith('[plugin]')) {
                if (!(detectLanguage in ttsPluginInfo.language)) {
                    throw new Error('Language not supported');
                }
                const config = (await store.get(serviceName)) ?? {};
                const data = await invoke('invoke_plugin', {
                    name: serviceName,
                    pluginType: 'tts',
                    source: sourceText,
                    lang: ttsPluginInfo.language[lang],
                    needs: config,
                });
                speak(data);
            } else {
                if (!(lang in builtinTtsServices[serviceName].Language)) {
                    throw new Error('Language not supported');
                }
                let data = await builtinTtsServices[serviceName].tts(
                    sourceText,
                    builtinTtsServices[serviceName].Language[lang]
                );
                speak(data);
            }
        } finally {
            setIsSpeakLoading(false)
        }
    };

    return (
        <Card
            shadow='none'
            className='bg-content1 rounded-[10px] mt-[1px] pb-0'
        >
            <Toaster />
            <CardBody className='bg-content1 p-[12px] pb-0 max-h-[40vh] overflow-y-auto'>
                <textarea
                    autoFocus
                    ref={textAreaRef}
                    className='bg-content1 h-full resize-none outline-none'
                    value={sourceText}
                    onKeyDown={keyDown}
                    onChange={(e) => {
                        const v = e.target.value;
                        setDetectLanguage('');
                        setSourceText(v);
                        if (dynamicTranslate) {
                            if (timer) {
                                clearTimeout(timer);
                            }
                            timer = setTimeout(() => {
                                detect_language(v).then(() => {
                                    syncSourceText();
                                });
                            }, 1000);
                        }
                    }}
                />
            </CardBody>
            <CardFooter className='bg-content1 rounded-none rounded-b-[10px] flex justify-between px-[12px] p-[5px]'>
                <div className='flex justify-start'>
                    <ButtonGroup className='mr-[5px]'>
                        <Tooltip content={t('translate.speak')}>
                            <Button
                                isIconOnly
                                variant='light'
                                size='sm'
                                onPress={() => {
                                    handleSpeak().catch((e) => {
                                        toast.error(e.toString(), { style: toastStyle });
                                    });
                                }}
                            >
                                {isSpeakLoading ? <Spinner size="sm" color="default" /> : <HiOutlineVolumeUp className='text-[16px]' />}
                            </Button>
                        </Tooltip>
                        <Tooltip content={t('translate.copy')}>
                            <Button
                                isIconOnly
                                variant='light'
                                size='sm'
                                onPress={() => {
                                    ClipboardSetText(sourceText).then((e) => {
                                        toast.success(e.toString(), { style: toastStyle });
                                    });
                                }}
                            >
                                <MdContentCopy className='text-[16px]' />
                            </Button>
                        </Tooltip>
                        <Tooltip content={t('translate.delete_newline')}>
                            <Button
                                isIconOnly
                                variant='light'
                                size='sm'
                                onPress={() => {
                                    const newText = sourceText.replace(/\s+/g, ' ');
                                    setSourceText(newText);
                                    detect_language(newText).then(() => {
                                        syncSourceText();
                                    });
                                }}
                            >
                                <MdSmartButton className='text-[16px]' />
                            </Button>
                        </Tooltip>
                        <Tooltip content={t('common.clear')}>
                            <Button
                                variant='light'
                                size='sm'
                                isIconOnly
                                isDisabled={sourceText === ''}
                                onPress={() => {
                                    setSourceText('');
                                }}
                            >
                                <LuDelete className='text-[16px]' />
                            </Button>
                        </Tooltip>
                    </ButtonGroup>
                    {detectLanguage !== '' && (
                        <Chip
                            size='sm'
                            color='secondary'
                            variant='dot'
                            className='my-auto'
                        >
                            {t(`languages.${detectLanguage}`)}
                        </Chip>
                    )}
                </div>
                <Button
                    size='sm'
                    color='primary'
                    variant='solid'
                    className='text-[14px] font-bold'
                    startContent={<HiTranslate className='text-[16px]' />}
                    onPress={() => {
                        console.log("resresresresr");
                        (async () => {
                            const worker = await createWorker('eng');
                            const ret = await worker.recognize('https://tesseract.projectnaptha.com/img/eng_bw.png');
                            console.log(ret.data.text);
                            console.log("resresresresr");
                            await worker.terminate();
                        })();
                        console.log("resresresresr");
                        // detect_language(sourceText).then(() => {
                        //     syncSourceText();
                        // });
                    }}
                >
                    {t('translate.translate')}
                </Button>
            </CardFooter>
        </Card>
    );
}
