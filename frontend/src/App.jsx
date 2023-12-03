import { useEffect } from 'react'
import { BrowserRouter } from 'react-router-dom';
import ToolBar from './window/ToolBar';
import Translate from './window/Translate';
import Screenshot from './window/Screenshot';
import { useTranslation } from 'react-i18next';
import './i18n';
import { useConfig, useSyncAtom } from './hooks';

const windowMap = {
  root: <ToolBar />,
  translate: <Translate />,
  screenshot: <Screenshot />,
};

function App({ variable }) {
  const { i18n } = useTranslation();
  const [appLanguage] = useConfig('app_language', 'zh_cn');

  useEffect(() => {
    console.log(variable, appLanguage)
    if (appLanguage !== null) {
      i18n.changeLanguage(appLanguage);
    }
  }, [appLanguage])

  return (
    <BrowserRouter>
      {windowMap[variable]}
    </BrowserRouter>
  )
}

export default App
