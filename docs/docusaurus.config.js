const config = require('./contrib/config.js')
const fs = require('fs')
const admonitions = require('remark-admonitions')

const githubRepoName =
  config.projectSlug === 'ecosystem' ? 'docs' : config.projectSlug

const baseUrl = config.baseUrl ? config.baseUrl : `/${config.projectSlug}/docs/`

const links = [
  {
    to: '/',
    activeBasePath: baseUrl,
    label: `Docs`,
    position: "left"
  },
  {
    href: "https://www.ory.sh/docs",
    label: "Ecosystem",
    position: "left"
  },
  {
    href: "https://www.ory.sh/blog", label: "Blog",
    position: "left"
  },
  {
    href: "https://community.ory.sh", label: "Forum",
    position: "left"
  },
  {
    href: "https://www.ory.sh/chat", label: "Chat",
    position: "left"
  },
  {
    href: `https://github.com/ory/${config.projectSlug}`,
    label: "GitHub",
    position: "left"
  }
];

let version = ["latest"];

if (fs.existsSync("./versions.json")) {
  version = require("./versions.json");
  if (version && version.length > 0) {
    links.push({
      label: version[0],
      position: "right",
      to: "versions"
    });
  }
  if (version.length === 0) {
    version = ["master"];
  }
}

const githubPrismTheme = require('prism-react-renderer/themes/github')

const prismThemeLight = {
  ...githubPrismTheme,
  styles: [
    ...githubPrismTheme.styles,
    {
      languages: ['keto-relation-tuples'],
      types: ['namespace'],
      style: {
        color: '#666'
      }
    },
    {
      languages: ['keto-relation-tuples'],
      types: ['object'],
      style: {
        color: '#939'
      }
    },
    {
      languages: ['keto-relation-tuples'],
      types: ['relation'],
      style: {
        color: '#e80'
      }
    },
    {
      languages: ['keto-relation-tuples'],
      types: ['delimiter'],
      style: {
        color: '#555'
      }
    },
    {
      languages: ['keto-relation-tuples'],
      types: ['comment'],
      style: {
        color: '#999'
      }
    },
    {
      languages: ['keto-relation-tuples'],
      types: ['subject'],
      style: {
        color: '#903'
      }
    }
  ]
}

module.exports = {
  title: config.projectName,
  tagline: config.projectTagLine,
  url: `https://www.ory.sh/`,
  baseUrl,
  favicon: 'img/favico.png',
  onBrokenLinks: 'error',
  onBrokenMarkdownLinks: 'error',
  organizationName: 'ory', // Usually your GitHub org/user name.
  projectName: config.projectSlug, // Usually your repo name.
  themeConfig: {
    prism: {
      theme: prismThemeLight,
      darkTheme: require('prism-react-renderer/themes/dracula'),
      additionalLanguages: ['pug', 'shell-session']
    },
    announcementBar: {
      id: 'supportus',
      content:
        config.projectSlug === 'docs'
          ? `Sign up for <a href="${config.newsletter}">important security announcements</a> and if you like the ${config.projectName} give us some ⭐️ on <a target="_blank" rel="noopener noreferrer" href="https://github.com/ory">GitHub</a>!`
          : `Sign up for <a href="${config.newsletter}">important security announcements</a> and if you like ${config.projectName} give it a ⭐️ on <a target="_blank" rel="noopener noreferrer" href="https://github.com/ory/${githubRepoName}">GitHub</a>!`
    },
    algolia: {
      apiKey: "8463c6ece843b377565726bb4ed325b0",
      indexName: "ory",
      algoliaOptions: {
        facetFilters: [`tags:${config.projectSlug}`, `version:${version[0]}`]
      }
    },
    navbar: {
      logo: {
        alt: config.projectName,
        src: `img/logo-${config.projectSlug}.svg`,
        srcDark: `img/logo-${config.projectSlug}.svg`,
        href:
          config.projectSlug === 'docs'
            ? `https://www.ory.sh`
            : `https://www.ory.sh/${config.projectSlug}`
      },
      items: links
    },
    footer: {
      style: 'dark',
      copyright: `Copyright © ${new Date().getFullYear()} ORY GmbH`,
      links: [
        {
          title: "Company",
          items: [
            {
              label: "Imprint",
              href: "https://www.ory.sh/imprint"
            },
            {
              label: "Privacy",
              href: "https://www.ory.sh/privacy"
            },
            {
              label: "Terms",
              href: "https://www.ory.sh/tos"
            }
          ]
        }
      ]
    }
  },
  plugins: [
    [
      "@docusaurus/plugin-content-docs",
      {
        path:
          config.projectSlug === 'docusaurus-template'
            ? 'contrib/docs'
            : 'docs',
        sidebarPath: require.resolve('./contrib/sidebar.js'),
        editUrl: ({ docPath }) =>
          `https://github.com/ory/${githubRepoName}/edit/master/docs/docs/${docPath}`,
        routeBasePath: '/',
        showLastUpdateAuthor: true,
        showLastUpdateTime: true,
        remarkPlugins: [admonitions]
      }
    ],
    "@docusaurus/plugin-content-pages",
    "@docusaurus/plugin-google-analytics",
    "@docusaurus/plugin-sitemap"
  ],
  themes: [
    [
      "@docusaurus/theme-classic",
      {
        customCss: config.projectSlug === "docusaurus-template" ? require.resolve("./contrib/theme.css") : require.resolve("./src/css/theme.css")
      }
    ],
    "@docusaurus/theme-search-algolia"
  ]
};
