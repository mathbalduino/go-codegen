const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

// With JSDoc @type annotations, IDEs can provide config autocompletion
/** @type {import('@docusaurus/types').DocusaurusConfig} */
(module.exports = {
  title: 'go-codegen',
  tagline: 'Everywhere, under @mathbalduino',
  url: 'https://mathbalduino.com.br',
  baseUrl: '/go-codegen/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',
  organizationName: 'mathbalduino', // Usually your GitHub org/user name.
  projectName: 'go-codegen', // Usually your repo name.
  trailingSlash: false,

  presets: [
    [
      '@docusaurus/preset-classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          editUrl: 'https://github.com/mathbalduino/go-codegen/edit/docs/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      colorMode: {
        defaultMode: 'dark'
      },
      navbar: {
        title: 'go-codegen',
        logo: {
          alt: '@mathbalduino logo',
          src: 'img/mathbalduino_logoM.png',
        },
        items: [
          {
            type: 'doc',
            docId: 'intro',
            position: 'left',
            label: 'Documentation',
          },
          {
            href: 'https://mathbalduino.com.br/about',
            position: 'left',
            label: 'About',
          },
          {
            href: 'https://github.com/mathbalduino/go-codegen',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          {
            title: 'Documentation',
            items: [
              {
                label: 'Introduction',
                to: '/docs/intro',
              },
            ],
          },
          {
            title: 'Author',
            items: [
              {
                label: 'mathbalduino.com.br',
                href: 'https://mathbalduino.com.br',
              },
            ],
          },
        ],
        logo: {
          alt: '@mathbalduino logo',
          href: 'http://mathbalduino.com.br',
          src: 'img/mathbalduino_logoS.png'
        },
        copyright: '@mathbalduino (Built with Docusaurus)',
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
    }),
});