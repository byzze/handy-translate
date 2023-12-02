import {
    Card,
    CardBody,
    CardHeader,
    CardFooter,
    Button,
    ButtonGroup,
    Skeleton,
    Dropdown,
    DropdownItem,
    DropdownMenu,
    DropdownTrigger,
    Tooltip,
} from '@nextui-org/react';
import toast, { Toaster } from 'react-hot-toast';
import React, { useEffect, useRef, useState } from "react";
import { useAtomValue } from 'jotai';
import { useTranslation } from 'react-i18next';
import { BiCollapseVertical, BiExpandVertical } from 'react-icons/bi';
import { TbTransformFilled } from 'react-icons/tb';
import { HiOutlineVolumeUp } from 'react-icons/hi';
import { MdContentCopy } from 'react-icons/md';
import { GiCycle } from 'react-icons/gi';

import * as builtinTtsServices from '../../../../services/tts';
import * as builtinServices from '../../../../services/translate';
import { useConfig, useToastStyle, useVoice } from '../../../../hooks';
import { sourceTextAtom, detectLanguageAtom } from '../SourceArea';
import { sourceLanguageAtom, targetLanguageAtom } from '../LanguageArea';


export default function TargetArea(props) {
    const { name, index, translateServiceList, ...drag } = props;

    const [collectionServiceList] = useConfig('collection_service_list', []);
    const [ttsServiceList] = useConfig('tts_service_list', ['lingva_tts']);
    const [autoCopy] = useConfig('translate_auto_copy', 'disable');
    const [clipboardMonitor] = useConfig('clipboard_monitor', false);
    const [hideWindow] = useConfig('translate_hide_window', false);

    const [hide, setHide] = useState(false);
    const [isSpeakLoading, setIsSpeakLoading] = useState(false)

    const [translateServiceName, setTranslateServiceName] = useState(name);
    const [isLoading, setIsLoading] = useState(false)
    const sourceText = useAtomValue(sourceTextAtom);
    const sourceLanguage = useAtomValue(sourceLanguageAtom);
    const targetLanguage = useAtomValue(targetLanguageAtom);
    const detectLanguage = useAtomValue(detectLanguageAtom);

    const speak = useVoice();
    const toastStyle = useToastStyle();

    const { t } = useTranslation();
    const textAreaRef = useRef();
    const [result, setResult] = useState('');
    const [error, setError] = useState('');

    useEffect(() => {
        // wails.Events.On("loading", function (data) {
        //     setIsLoading(data.data == 'true')
        // })
        wails.Events.On("result", function (data) {
            let result = data.data
            setResult(result)
        })
        const LanguageEnum = builtinServices[translateServiceName].Language;
        if (sourceLanguage in LanguageEnum && targetLanguage in LanguageEnum) {
            wails.Events.Emit({ name: "translateLang", data: [LanguageEnum[sourceLanguage], LanguageEnum[targetLanguage]] })
        }
    }, [targetLanguage, sourceLanguage])

    useEffect(() => {
        setResult('');
        setError('');

        if (sourceText !== '' && sourceLanguage && targetLanguage) {
            const LanguageEnum = builtinServices[translateServiceName].Language;

            if (sourceLanguage in LanguageEnum && targetLanguage in LanguageEnum) {
                window.go.main.App.Transalte(sourceText, LanguageEnum[sourceLanguage], LanguageEnum[targetLanguage])
            }
        }
    }, [sourceText, targetLanguage, sourceLanguage, autoCopy, hideWindow, translateServiceName, clipboardMonitor]);

    const handleSpeak = async () => {
        try {
            setIsSpeakLoading(true)
            const serviceName = ttsServiceList[0];
            let data = await builtinTtsServices[serviceName].tts(
                result,
                builtinTtsServices[serviceName].Language[targetLanguage]
            );

            speak(data);
        } finally {
            setIsSpeakLoading(false)
        }
    };

    useEffect(() => {
        if (textAreaRef.current !== null) {
            textAreaRef.current.style.height = '0px';
            if (result !== '') {
                textAreaRef.current.style.height = textAreaRef.current.scrollHeight + 'px';
            }
        }
    }, [result]);

    return (
        <Card
            shadow='none'
            className='rounded-[10px]'
        >
            <Toaster />
            <CardHeader
                className={`flex justify-between py-1 px-0 bg-content2 h-[30px] ${hide ? 'rounded-[10px]' : 'rounded-t-[10px]'
                    }`}
                {...drag}
            >
                <div className='flex'>
                    <Dropdown>
                        <DropdownTrigger>
                            <Button
                                size='sm'
                                variant='solid'
                                className='bg-transparent'
                                startContent={
                                    translateServiceName.startsWith('[plugin]') ? (
                                        <img
                                            src={[translateServiceName].icon}
                                            className='h-[20px] my-auto'
                                        />
                                    ) : (
                                        <img
                                            src={builtinServices[translateServiceName].info.icon}
                                            className='h-[20px] my-auto'
                                        />
                                    )
                                }
                            >
                                {translateServiceName.startsWith('[plugin]') ? (
                                    <div className='my-auto'>{`${[translateServiceName].display} `}</div>
                                ) : (
                                    <div className='my-auto'>
                                        {t(`services.translate.${translateServiceName}.title`)}
                                    </div>
                                )}
                            </Button>
                        </DropdownTrigger>
                        <DropdownMenu
                            aria-label='app language'
                            className='max-h-[40vh] overflow-y-auto'
                            onAction={(key) => {
                                setTranslateServiceName(key);
                            }}
                        >
                            {translateServiceList.map((x) => {
                                return (
                                    <DropdownItem
                                        key={x}
                                        startContent={
                                            x.startsWith('[plugin]') ? (
                                                <img
                                                    src={[x].icon}
                                                    className='h-[20px] my-auto'
                                                />
                                            ) : (
                                                <img
                                                    src={builtinServices[x].info.icon}
                                                    className='h-[20px] my-auto'
                                                />
                                            )
                                        }
                                    >
                                        {x.startsWith('[plugin]') ? (
                                            <div className='my-auto'>{`${[x].display} `}</div>
                                        ) : (
                                            <div className='my-auto'>{t(`services.translate.${x}.title`)}</div>
                                        )}
                                    </DropdownItem>
                                );
                            })}
                        </DropdownMenu>
                    </Dropdown>
                </div>
                <div className='flex'>
                    <Button
                        size='sm'
                        isIconOnly
                        variant='light'
                        className='h-[20px] w-[20px]'
                        onPress={() => setHide(!hide)}
                    >
                        {hide ? (
                            <BiExpandVertical className='text-[16px]' />
                        ) : (
                            <BiCollapseVertical className='text-[16px]' />
                        )}
                    </Button>
                </div>
            </CardHeader>
            <CardBody
                className={`p-[12px] pb-0 ${hide && 'hidden'} ${result === '' && error === '' && !isLoading && 'hidden'
                    }`}
            >
                {isLoading ? (
                    <div className='space-y-3'>
                        <Skeleton className='w-4/5 rounded-lg'>
                            <div className='h-3 w-4/5 rounded-lg bg-default-200'></div>
                        </Skeleton>
                        <Skeleton className='w-3/5 rounded-lg'>
                            <div className='h-3 w-3/5 rounded-lg bg-default-200'></div>
                        </Skeleton>
                    </div>
                ) : typeof result === 'string' ? (
                    <textarea
                        ref={textAreaRef}
                        className='h-0 resize-none bg-transparent select-text outline-none'
                        readOnly
                        value={result}
                    />
                ) : (
                    <div>
                        {result['pronunciations'] &&
                            result['pronunciations'].map((pronunciation) => {
                                return (
                                    <div key={nanoid()}>
                                        {pronunciation['region'] && (
                                            <span className='mr-[12px] text-default-500'>
                                                {pronunciation['region']}
                                            </span>
                                        )}
                                        {pronunciation['symbol'] && (
                                            <span className='mr-[12px] text-default-500'>
                                                {pronunciation['symbol']}
                                            </span>
                                        )}
                                        {pronunciation['voice'] && pronunciation['voice'] !== '' && (
                                            <HiOutlineVolumeUp
                                                className='inline-block my-auto mr-[12px] cursor-pointer'
                                                onClick={() => {
                                                    speak(pronunciation['voice']);
                                                }}
                                            />
                                        )}
                                    </div>
                                );
                            })}
                        {result['explanations'] &&
                            result['explanations'].map((explanations) => {
                                return (
                                    <div key={nanoid()}>
                                        {explanations['explains'] &&
                                            explanations['explains'].map((explain, index) => {
                                                return (
                                                    <span key={nanoid()}>
                                                        {index === 0 ? (
                                                            <>
                                                                <span className='text-[12px] text-default-500 mr-[12px]'>
                                                                    {explanations['trait']}
                                                                </span>
                                                                <span className='font-bold text-[16px] select-text'>
                                                                    {explain}
                                                                </span>
                                                                <br />
                                                            </>
                                                        ) : (
                                                            <span
                                                                className='text-[14px] text-default-500 mr-[8px] select-text'
                                                                key={nanoid()}
                                                            >
                                                                {explain}
                                                            </span>
                                                        )}
                                                    </span>
                                                );
                                            })}
                                    </div>
                                );
                            })}
                        <br />
                        {result['associations'] &&
                            result['associations'].map((association) => {
                                return (
                                    <div key={nanoid()}>
                                        <span className='mr-[12px] text-default-500'>{association}</span>
                                    </div>
                                );
                            })}
                        {result['sentence'] &&
                            result['sentence'].map((sentence, index) => {
                                return (
                                    <div key={nanoid()}>
                                        <span className='mr-[12px]'>{index + 1}.</span>
                                        <>
                                            {sentence['source'] && (
                                                <span
                                                    className='select-text'
                                                    dangerouslySetInnerHTML={{
                                                        __html: sentence['source'],
                                                    }}
                                                />
                                            )}
                                        </>
                                        <>
                                            {sentence['target'] && (
                                                <div
                                                    className='select-text text-default-500'
                                                    dangerouslySetInnerHTML={{
                                                        __html: sentence['target'],
                                                    }}
                                                />
                                            )}
                                        </>
                                    </div>
                                );
                            })}
                    </div>
                )}
                {error !== '' ? (
                    error.split('\n').map((v) => {
                        return (
                            <p
                                key={v}
                                className='text-red-500'
                            >
                                {v}
                            </p>
                        );
                    })
                ) : (
                    <></>
                )}
            </CardBody>
            <CardFooter
                className={`bg-content1 rounded-none rounded-b-[10px] flex px-[12px] p-[5px] ${hide && 'hidden'}`}
            >
                <ButtonGroup>
                    <Tooltip content={t('translate.speak')}>
                        <Button
                            isIconOnly
                            variant='light'
                            size='sm'
                            isDisabled={typeof result !== 'string' || result === ''}
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
                            isDisabled={typeof result !== 'string' || result === ''}
                            onPress={() => {
                                wails.Clipboard.SetText(result).then((e) => {
                                    toast.success(e.toString(), { style: toastStyle });
                                });
                            }}
                        >
                            <MdContentCopy className='text-[16px]' />
                        </Button>
                    </Tooltip>
                    <Tooltip content={t('translate.retry')}>
                        <Button
                            isIconOnly
                            variant='light'
                            size='sm'
                            className={`${error === '' && 'hidden'}`}
                            onPress={() => {
                                setError('');
                                setResult('');
                                translate();
                            }}
                        >
                            <GiCycle className='text-[16px]' />
                        </Button>
                    </Tooltip>
                    {collectionServiceList &&
                        collectionServiceList.map((serviceName) => {
                            return (
                                <Button
                                    key={serviceName}
                                    isIconOnly
                                    variant='light'
                                    size='sm'
                                    onPress={async () => {
                                        if (serviceName.startsWith('[plugin]')) {
                                            const pluginConfig = (await store.get(serviceName)) ?? {};
                                            invoke('invoke_plugin', {
                                                name: serviceName,
                                                pluginType: 'collection',
                                                source: sourceText.trim(),
                                                target: result.toString(),
                                                from: detectLanguage,
                                                to: targetLanguage,
                                                needs: pluginConfig,
                                            }).then(
                                                (_) => {
                                                    toast.success(t('translate.add_collection_success'), {
                                                        style: toastStyle,
                                                    });
                                                },
                                                (e) => {
                                                    toast.error(e.toString(), { style: toastStyle });
                                                }
                                            );
                                        } else {
                                            builtinCollectionServices[serviceName].collection(sourceText, result).then(
                                                (_) => {
                                                    toast.success(t('translate.add_collection_success'), {
                                                        style: toastStyle,
                                                    });
                                                },
                                                (e) => {
                                                    toast.error(e.toString(), { style: toastStyle });
                                                }
                                            );
                                        }
                                    }}
                                >
                                    <img
                                        src={
                                            serviceName.startsWith('[plugin]')
                                                ? [serviceName].icon
                                                : builtinCollectionServices[serviceName].info.icon
                                        }
                                        className='h-[16px] w-[16px]'
                                    />
                                </Button>
                            );
                        })}
                </ButtonGroup>
            </CardFooter>
        </Card>
    );
}
