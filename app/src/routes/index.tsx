import { createSignal, For, Show } from 'solid-js';
import { Categorie, CategorieItem } from '../types/index';
import { CATEGORIES_MOCK } from '~/mockData';
import SideMenu from '~/components/SideMenu';
import ActiveVot from '~/components/ActiveVot';

export default function Home() {
  const [activeCategorie, setActiveCategorie] = createSignal<Categorie>(
    CATEGORIES_MOCK[0],
    {
      equals: false,
    }
  );
  const [activeVot, setActiveVot] = createSignal<CategorieItem>();
  const [categories, setCategories] =
    createSignal<Categorie[]>(CATEGORIES_MOCK);

  return (
    <main class="text-gray-70 w-full flex">
      <SideMenu
        categories={categories}
        handleActiveCategoriesItems={setActiveCategorie}
        activeCategorie={activeCategorie}
        handleItem={setActiveVot}
      />
      <section class="p-2 w-full border">
        <Show when={activeVot()} fallback={<p>Please select an item</p>}>
          <ActiveVot activeVot={activeVot} />
        </Show>
      </section>
    </main>
  );
}
