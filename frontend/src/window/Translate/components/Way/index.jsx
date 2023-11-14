import React from "react";
import { RadioGroup, Radio } from "@nextui-org/react";
import { useEffect } from "react";
import toast, { Toaster } from 'react-hot-toast';
import { GetTransalteMap, GetTransalteWay, SetTransalteWay } from "../../../../../wailsjs/go/main/App"
import { useConfig, useToastStyle, useVoice, useSyncAtom } from '../../../../hooks';
import { atom, useAtom, useAtomValue } from 'jotai';

export const translateServiceListAtom = atom([]);

let timer = null;
export default function Way() {
    const [translateMap, setTranslateMap] = React.useState({});

    const [translateServiceList, setTranslateServiceList, syncTranslateServiceList] = useSyncAtom(translateServiceListAtom)

    const [selected, setSelected] = React.useState("");
    const toastStyle = useToastStyle();

    useEffect(() => {
        // GetTransalteWay().then(result => {
        //     setTranslateServiceList([result]);
        // });

        GetTransalteMap().then(result => {
            result = JSON.parse(result)
            setTranslateMap(result)
        })

        GetTransalteWay().then(result => {
            setSelected(result)
        })
    }, [])


    return (
        <div className="flex flex-col gap-3">
            <Toaster />
            <RadioGroup
                label="选择你想要的翻译服务"
                value={selected}
                onValueChange={(value => {
                    console.log(value)
                    setSelected(value)
                    SetTransalteWay(value)
                    setTranslateServiceList([value])

                    if (timer) {
                        clearTimeout(timer);
                    }

                    timer = setTimeout(() => {
                        syncTranslateServiceList()
                    }, 100);

                })}
            >

                {translateMap &&
                    Object.entries(translateMap).map((key) => {
                        return (<Radio key={key[0]} value={key[0]}>{key[1].name}</Radio>)
                    })
                }
            </RadioGroup>
            {/* <p className="text-default-500 text-small">Selected: {selected}</p> */}
        </div>
    );
}
