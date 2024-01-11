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

  const leftButtons = document.querySelector('.left-buttons');
  const menuBar = document.querySelector('.menu-bar');
  const searchButton = document.createElement('button');
  searchButton.type = 'button';
  searchButton.classList.add('icon-button');
  searchButton.title = 'Пошук. (Швидка клавіша: ш)';
  searchButton.setAttribute('aria-label', 'Показати пошук');
  searchButton.setAttribute('aria-expanded', 'false');
  searchButton.setAttribute('aria-keyshortcuts', 'ш');
  searchButton.setAttribute('aria-controls', 'searchbar');
  const searchIcon = document.createElement('i');
  searchIcon.classList.add('fa', 'fa-search');
  searchButton.append(searchIcon);
  leftButtons.append(searchButton);
  const searchWrapper = document.createElement('div');
  searchWrapper.id = 'search-wrapper';
  const searchBox = document.createElement('div');
  searchBox.id = 'searchbar-outer';
  searchBox.classList.add('searchbar-outer');
  searchWrapper.classList.add('hidden');
  searchWrapper.append(searchBox);

  const searchresultsOuter = document.createElement('div');
  searchresultsOuter.id = 'searchresults-outer';
  searchresultsOuter.classList.add('searchresults-outer');
  searchWrapper.append(searchresultsOuter);

  const hits = document.createElement('div');
  hits.id = 'hits';
  searchresultsOuter.append(hits);
  const closeButton = document.createElement('button');
  closeButton.classList.add('close-button');
  menuBar.parentElement.insertBefore(searchWrapper, menuBar.nextElementSibling);
  //   leftButtons.append(searchBox, hits, closeButton);

  function activateSearch() {
    searchButton.setAttribute('aria-expanded', 'true');
    searchButton.setAttribute('aria-label', 'Сховати пошук');
    // searchBox.classList.add('active');
    searchWrapper.classList.remove('hidden');
    hits.classList.add('active');
    document.querySelector('.ais-SearchBox-input').focus();
    // document.body.style.overflow = 'hidden';
  }

  function deactivateSearch() {
    searchButton.setAttribute('aria-expanded', 'false');
    searchButton.setAttribute('aria-label', 'Показати пошук');
    // searchBox.classList.remove('active');
    searchWrapper.classList.add('hidden');
    hits.classList.remove('active');
    // document.body.style.overflow = '';
  }

  function triggerSearch(event) {
    if (event.currentTarget.getAttribute('aria-expanded') === 'true') {
      deactivateSearch();
    } else {
      activateSearch();
    }
  }
  window.addEventListener('keydown', (event) => {
    if (event.key === 'i') {
      //   event.preventDefault();
      triggerSearch(event);
    }
  });

  searchButton.addEventListener('click', triggerSearch);
  closeButton.addEventListener('click', (event) => {
    event.stopPropagation();
    deactivateSearch();
  });

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
    searchFunction(helper) {
      searchresultsOuter.style.display = helper.state.query === '' ? 'none' : '';

      helper.search();
    },
  });

  search.addWidgets([
    window.instantsearch.widgets.searchBox({
      container: searchBox,
      cssClasses: {
        form: 'searchbar-outer',
      },
    }),

    window.instantsearch.widgets.hits({
      container: hits,
      templates: {
        empty: 'Нічого не знайдено на тему {{query}}.',
        item(hit, { html, components }) {
          return html`
                    <a href="/uk/docs/${hit.slug}">${components.Highlight({ hit, attribute: 'title' })}</a>
                    ${components.Snippet({ classNames: { root: 'teaser' }, hit, attribute: 'text' })}
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
