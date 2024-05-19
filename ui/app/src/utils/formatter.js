const cachedFormatters = {};

const getLanguage = () => {
    if (window.navigator.languages) {
        return window.navigator.languages[0];
    } else {
        return window.navigator.language;
    }
}

export const formatCurrency = (value, currency) => {
    if (cachedFormatters[currency] === undefined) {
        const language = getLanguage()
        cachedFormatters[currency]
            = new Intl.NumberFormat(language, {style: 'currency', currency: currency});
    }
    
    return cachedFormatters[currency].format(value);
}
