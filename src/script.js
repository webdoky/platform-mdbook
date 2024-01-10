document.addEventListener('DOMContentLoaded', () => {
  const missingLinks = document.querySelectorAll('.missing-link');
  missingLinks.forEach((link) => {
    link.addEventListener('click', (event) => {
      event.preventDefault();
      const userResponse = confirm('Ми не маємо перекладу цієї сторінки. Спробувати відкрити варіант англійською?');
      if (userResponse) {
        const url = new URL(event.currentTarget.href);
        let targetPathname = url.pathname;
        targetPathname = `/en-US${targetPathname.slice('/uk'.length)}`;
        window.open(`https://developer.mozilla.org${targetPathname}`, '_blank');
      }
    });
  });

  //   const nolintCodeBlocks =
  // document.querySelectorAll('code[class*="language-"][class*="-nolint"]');
  //   nolintCodeBlocks.forEach((codeBlock) => {
  //     codeBlock.classList.forEach((className) => {
  //       if (className.endsWith('-nolint')) {
  //         codeBlock.classList.remove(className);
  //         codeBlock.classList.add(className.slice(0, -'-nolint'.length));
  //       }
  //     });
  //   });

  const leftButtons = document.querySelector('.left-buttons');
  const searchBox = document.createElement('div');
  searchBox.id = 'searchbox';
  const hits = document.createElement('div');
  hits.id = 'hits';
  const closeButton = document.createElement('button');
  closeButton.classList.add('close-button');
  leftButtons.append(searchBox, hits, closeButton);

  function activateSearch() {
    searchBox.classList.add('active');
    hits.classList.add('active');
    document.body.style.overflow = 'hidden';
  }

  function deactivateSearch() {
    searchBox.classList.remove('active');
    hits.classList.remove('active');
    document.body.style.overflow = '';
  }

  function handleClick(event) {
    // if clicked outside of the search box
    if (!event.target.closest('#searchbox') && !event.target.closest('.hits')) {
      // hide the search box
      deactivateSearch();
    } else {
      // show the search box
      activateSearch();
    }
  }
  closeButton.addEventListener('click', (event) => {
    event.stopPropagation();
    deactivateSearch();
  });

  document.addEventListener('click', handleClick);

  document.addEventListener('keydown', (event) => {
    if (event.key === 'Escape') {
      deactivateSearch();
    }
  });

  const searchClient = window.algoliasearch('CYCK8RSV1M', '174e3933e9859ac40e716172c477daee');

  const search = window.instantsearch({
    indexName: 'articles',
    insights: true,
    searchClient,
  });

  search.addWidgets([
    window.instantsearch.widgets.searchBox({
      container: '#searchbox',
    }),

    window.instantsearch.widgets.hits({
      container: '#hits',
      templates: {
        item(hit, { html, components }) {
          return html`
            <a href="/uk/docs/${hit.slug}"><h2>${components.Highlight({ hit, attribute: 'title' })}</h2>
            <p>${components.Snippet({ hit, attribute: 'slug' })}</p></a>
          `;
        },
      },
    }),
  ]);

  search.start();

  document.querySelectorAll('.hljs.example-bad').forEach((element) => {
    element.parentElement.querySelector('.buttons .clip-button').classList.add('hidden');
  });
});
