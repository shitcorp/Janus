/**
 * Checks if string has a number
 * @param str
 */
export const hasNumber = (str: string): boolean => {
  return /\d/.test(str);
};
