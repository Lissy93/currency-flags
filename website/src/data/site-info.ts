
// Repository info (makes is easier to re-use)

const basics = {
  title: 'Currency Flags',
  author: 'Alicia Sykes',
  description:
    `A collection of flag icons, associated with world currencies.
    Consumable via a REST API, CDN, or NPM package.`,
};

const repoInfo = {
  srcProvider: 'https://github.com',  // The source control provider
  srcUsername: 'lissy93',             // Username associated with repo
  srcRepoName: 'currency-flags',      // Repository name, case-sensitive
  srcBranch: 'master',                // The default branch
};

const repo = `${repoInfo.srcProvider}/${repoInfo.srcUsername}/${repoInfo.srcRepoName}`;

const navigation = [
  {
    title: 'API',
    path: '/api',
    description: 'Fetch currency info and associated flag icons via the REST API.',
  },
  {
    title: 'CDN',
    path: '/cdn',
    description: 'Directly embbed currency flag icons, sourced from an edge CDN.',
  },
  {
    title: 'CSS',
    path: '/css',
    description: 'Use the NPM-installable stylesheet to render currency flag assets as CSS.',
  },
  {
    title: 'Flags',
    path: '/#supported-currencies',
    description: 'View list of supported currencies.',
  }
];

const footer = {
  project: {
    name: basics.title,
    link: repo,
  },
  author: {
    name: basics.author,
    link: `${repoInfo.srcProvider}/${repoInfo.srcUsername}`,
  },
  license: {
    name: 'MIT',
    link: `${repo}/blob/${repoInfo.srcBranch}/LICENSE`,
  },
}

export default {
  ...basics,
  navigation,
  footer,
};

