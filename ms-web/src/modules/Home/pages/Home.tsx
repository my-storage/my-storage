import { createSignal } from 'solid-js';

import { Box } from '@components/Box';

export function HomePage() {
  const [value, setValue] = createSignal(true);

  return (
    <section>
      <button onClick={() => setValue(!value())}>click</button>
      <div>divider</div>
      <div style={{ background: value() ? 'red' : 'blue' }}>Foo</div>
      <div>divider</div>
      <Box
        display="flex"
        alignItems="center"
        justifyContent="center"
        flexDirection="column"
        background={value() ? 'red' : 'blue'}
        fontWeight={600}
        class="foo"
      >
        {JSON.stringify(value())}
      </Box>
    </section>
  );
}
