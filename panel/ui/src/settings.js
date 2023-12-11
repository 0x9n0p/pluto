export const Address = () => {
  let host = getValueFromLocalStorage('base_host');
  if (host === null)
    return '';
  else
    return 'https://' + host;
};

export const getValueFromLocalStorage = (name) => {
  if (typeof window !== 'undefined')
    return localStorage.getItem(name);
  else
    return null;
};
