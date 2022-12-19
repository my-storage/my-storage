/* eslint-disable @typescript-eslint/no-explicit-any */

import { CSSProperties, cssPropertiesKeys } from './types';

export function createStyleFromProps(values: any): CSSProperties {
  const result: any = {};

  for (const p of Object.keys(values)) {
    const prop = p as keyof CSSProperties;

    if (cssPropertiesKeys.includes(prop)) {
      if (values[prop] !== undefined) {
        result[prop] = values[prop];
      }
    }
  }

  return result;
}
