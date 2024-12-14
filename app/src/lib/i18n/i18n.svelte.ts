import en from './translations.en'

export interface Language {
    [key: string]: string
}

export type Translation = {
    [key in keyof typeof en]: string;
};

interface I18n {
    [locale: string]: Language
}

const defaultLocale = "en"
const supportedLocales = ["de", "en"]
const translations: I18n = {en}

const getInitialLocale = () => {
    if (navigator.language && supportedLocales.indexOf(navigator.language) >= 0) {
        return navigator.language
    }
    return defaultLocale
}

let currentLocale = $state(getInitialLocale())

export const localeSwitcher = {
    get locale() {
        return currentLocale
    },
    set locale(value) {
        if (!(value in translations)) {
            import(`./translations.${value}.ts`)
                .then(file => translations[value] = file.default)
                .then(() => currentLocale = value)
        } else {
            currentLocale = value
        }
    }
}

const translate = (locale: string, key: keyof Translation, vars?: Record<string, any>) => {
    if (!key) throw new Error("no key provided to $t()");
    if (!locale) throw new Error(`no translation for key "${key}"`);

    let text = translations[locale][key];

    if (!text) throw new Error(`no translation found for ${locale}.${key}`);

    if (vars) {
        Object.keys(vars).map((k) => {
            const regex = new RegExp(`{{${k}}}`, "g");
            text = text.replace(regex, vars[k]);
        });
    }

    return text;
}

export default (key: keyof Translation, vars?: any) => {
    return translate(currentLocale, key, vars)
}