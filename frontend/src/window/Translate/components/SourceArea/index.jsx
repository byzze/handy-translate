import React, { useEffect, useRef, useState } from "react";
import { Button, Card, CardBody, CardFooter, ButtonGroup, Chip, Tooltip } from '@nextui-org/react';
import toast, { Toaster } from 'react-hot-toast';
import { useSyncAtom } from '../../../../hooks';
import { atom, useAtom } from 'jotai';
import { useTranslation } from 'react-i18next';
import { HiOutlineVolumeUp } from 'react-icons/hi';
import { HiTranslate } from 'react-icons/hi';
import { LuDelete } from 'react-icons/lu';
import { MdContentCopy } from 'react-icons/md';
import { MdSmartButton } from 'react-icons/md';
import { WindowHide, EventsOn } from "../../../../../wailsjs/runtime"

export const sourceTextAtom = atom('sourceTextAtom');
export const detectLanguageAtom = atom('auto');

export default function SourceArea(props) {
    const { pluginList } = props;
    const [sourceText, setSourceText, syncSourceText] = useSyncAtom(sourceTextAtom);
    const [detectLanguage, setDetectLanguage] = useAtom(detectLanguageAtom);
    // const [recognizeLanguage] = useConfig('recognize_language', 'auto');
    const textAreaRef = useRef();
    const { t } = useTranslation();

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
    }, [sourceText]);

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
                                <HiOutlineVolumeUp className='text-[16px]' />
                            </Button>
                        </Tooltip>
                        <Tooltip content={t('translate.copy')}>
                            <Button
                                isIconOnly
                                variant='light'
                                size='sm'
                                onPress={() => {
                                    writeText(sourceText);
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
                        detect_language(sourceText).then(() => {
                            syncSourceText();
                        });
                    }}
                >
                    {t('translate.translate')}
                </Button>
            </CardFooter>
        </Card>
    );
}
