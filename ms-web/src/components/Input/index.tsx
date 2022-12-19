import { splitProps } from 'solid-js';
import { css, StylesArg } from 'solid-styled-components';

import { StyledDynamic } from '@features/StyledDynamic';

import { CSSProperties } from '@shared/utils/css-properties';

import type { ComponentWithAs } from '@shared/types/component-with-as';

export const Input: ComponentWithAs<CSSProperties> = (_props) => {
  const [props, rest] = splitProps(_props, ['as', 'children']);

  return (
    <StyledDynamic as={props.as ?? 'input'} {...rest}>
      {props.children}
    </StyledDynamic>
  );
};
