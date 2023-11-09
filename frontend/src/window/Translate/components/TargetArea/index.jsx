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

import { sourceLanguageAtom, targetLanguageAtom } from '../LanguageArea';
import { sourceTextAtom, detectLanguageAtom } from '../SourceArea';
import { useConfig, useToastStyle, useVoice } from '../../../../hooks';
import { WindowHide, EventsOn, ClipboardSetText } from "../../../../../wailsjs/runtime"
import * as builtinTtsServices from '../../../../services/tts';
import * as builtinServices from '../../../../services/translate';



export default function TargetArea(props) {
    const [collectionServiceList] = useConfig('collection_service_list', []);
    const [ttsServiceList] = useConfig('tts_service_list', ['lingva_tts']);
    const { name, index, translateServiceList, pluginList, ...drag } = props;
    const [hide, setHide] = useState(false);
    const [autoCopy] = useConfig('translate_auto_copy', 'disable');
    const [translateServiceName, setTranslateServiceName] = useState(name);
    const [clipboardMonitor] = useConfig('clipboard_monitor', false);
    const [hideWindow] = useConfig('translate_hide_window', false);
    const [isLoading, setIsLoading, getIsLoading] = useState(false);
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
        setResult('');
        setError('');
        EventsOn("loading", (result) => {
            setIsLoading(result);
        })
        if (
            sourceText.trim() !== '' &&
            sourceLanguage &&
            targetLanguage &&
            autoCopy !== null &&
            hideWindow !== null &&
            clipboardMonitor !== null
        ) {
            if (autoCopy === 'source' && !clipboardMonitor) {
                ClipboardSetText(sourceText).then((e) => {
                    toast.success(e.toString(), { style: toastStyle });
                    if (hideWindow) {
                        sendNotification({ title: t('common.write_clipboard'), body: sourceText });
                    }
                });
            }

            // setIsLoading(true)

            // translate();
            EventsOn("result", (result) => {
                setResult(result)
                setIsLoading(false)
            })
        }
    }, [sourceText, targetLanguage, sourceLanguage, autoCopy, hideWindow, translateServiceName, clipboardMonitor]);

    const handleSpeak = async () => {
        const serviceName = ttsServiceList[0];

        if (serviceName.startsWith('[plugin]')) {
            const config = (await store.get(serviceName)) ?? {};
            if (!(targetLanguage in ttsPluginInfo.language)) {
                throw new Error('Language not supported');
            }
            let data = await invoke('invoke_plugin', {
                name: serviceName,
                pluginType: 'tts',
                source: result,
                lang: ttsPluginInfo.language[targetLanguage],
                needs: config,
            });
            speak(data);
        } else {
            if (!(targetLanguage in builtinTtsServices[serviceName].Language)) {
                throw new Error('Language not supported');
            }
            let data = await builtinTtsServices[serviceName].tts(
                result,
                builtinTtsServices[serviceName].Language[targetLanguage]
            );
            speak(data);
        }
    };

    useEffect(() => {
        if (ttsServiceList && ttsServiceList[0].startsWith('[plugin]')) {
            readTextFile(`plugins/tts/${ttsServiceList[0]}/info.json`, {
                dir: BaseDirectory.AppConfig,
            }).then((infoStr) => {
                setTtsPluginInfo(JSON.parse(infoStr));
            });
        }
    }, [ttsServiceList]);

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
                                            src={pluginList['translate'][translateServiceName].icon}
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
                                    <div className='my-auto'>{`${pluginList['translate'][translateServiceName].display} `}</div>
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
                                                    src={pluginList['translate'][x].icon}
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
                                            <div className='my-auto'>{`${pluginList['translate'][x].display} `}</div>
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
                            <HiOutlineVolumeUp className='text-[16px]' />
                        </Button>
                    </Tooltip>
                    <Tooltip content={t('translate.copy')}>
                        <Button
                            isIconOnly
                            variant='light'
                            size='sm'
                            isDisabled={typeof result !== 'string' || result === ''}
                            onPress={() => {
                                ClipboardSetText(result).then((e) => {
                                    toast.success(e.toString(), { style: toastStyle });
                                });
                            }}
                        >
                            <MdContentCopy className='text-[16px]' />
                        </Button>
                    </Tooltip>
                    <Tooltip content={t('translate.translate_back')}>
                        <Button
                            isIconOnly
                            variant='light'
                            size='sm'
                            isDisabled={typeof result !== 'string' || result === ''}
                            onPress={async () => {

                            }}
                        >
                            <TbTransformFilled className='text-[16px]' />
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
                                                ? pluginList['collection'][serviceName].icon
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
