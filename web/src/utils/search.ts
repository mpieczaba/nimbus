export const extractSearchOperators = (
  searchString: string | null
): string | null => {
  if (typeof searchString !== "string" || searchString.length === 0) {
    return null;
  }

  if (searchString.match(/".+"/))
    return searchString.substring(1, searchString.length - 1);

  return `%${searchString}%`;
};
