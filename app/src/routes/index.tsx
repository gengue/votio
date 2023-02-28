import { createSignal, For, Show } from "solid-js";
import { createStore } from "solid-js/store";

export default function Home() {
	const [activeCategories, setActiveCategories] = createSignal([], {
		equals: false,
	});
	const [activeItem, setActiveItem] = createSignal(null, {
		equals: false,
	});
	const [categories, setCategories] = createSignal([
		{
			categorie: "Website",
			items: [
				{
					id: "001",
					votes: 2,
					title: "Test title 1",
					description:
						"loremlAute id cillum anim adipisicing labore laboris commodo consequat esse excepteur adipisicing. Tempor consequat quis aliquip aliqua ipsum et sint cupidatat consequat. Ipsum excepteur ullamco nostrud exercitation laboris sit dolor in consectetur officia.",
				},
				{
					id: "002",
					votes: 1,
					title: "Test title 2",
					description:
						"loremlAute id cillum anim adipisicing labore laboris commodo consequat esse excepteur adipisicing. Tempor consequat quis aliquip aliqua ipsum et sint cupidatat consequat. Ipsum excepteur ullamco nostrud exercitation laboris sit dolor in consectetur officia.",
				},
			],
		},
		{
			categorie: "Go",
			items: [
				{
					id: "001",
					votes: 2,
					title: "Test title 1",
					description:
						"loremlFugiat proident nostrud commodo mollit esse. Irure fugiat enim consequat proident eu eiusmod in ipsum duis nisi enim nulla ut irure. In ipsum voluptate est cupidatat consequat elit occaecat consectetur. Esse nulla laboris incididunt ullamco ullamco consectetur laborum aliquip nisi.",
				},
				{
					id: "002",
					votes: 1,
					title: "Test title 2",
					description:
						"loremlAute id cillum anim adipisicing labore laboris commodo consequat esse excepteur adipisicing. Tempor consequat quis aliquip aliqua ipsum et sint cupidatat consequat. Ipsum excepteur ullamco nostrud exercitation laboris sit dolor in consectetur officia.",
				},
			],
		},
		{
			categorie: "Data",
			items: [
				{
					id: "001",
					votes: 2,
					title: "Test title 1",
					description:
						"loremlSit laboris adipisicing mollit ex pariatur in excepteur adipisicing. Consequat quis mollit tempor aliqua ipsum est laboris do cupidatat excepteur anim elit in proident. Velit sunt in velit magna.",
				},
				{
					id: "002",
					votes: 1,
					title: "Test title 2",
					description:
						"loremlAute id cillum anim adipisicing labore laboris commodo consequat esse excepteur adipisicing. Tempor consequat quis aliquip aliqua ipsum et sint cupidatat consequat. Ipsum excepteur ullamco nostrud exercitation laboris sit dolor in consectetur officia.",
				},
			],
		},
		{
			categorie: "Hubspot",
			items: [
				{
					id: "001",
					votes: 2,
					title: "Test title 1",
					description:
						"loremlExcepteur tempor consectetur aute in qui cupidatat irure adipisicing officia id enim. Dolore ipsum ut exercitation voluptate ex voluptate id. Duis nostrud eiusmod eu nostrud proident fugiat aliqua ex nisi dolor irure. Quis ex laboris cillum eiusmod ad culpa nisi in excepteur exercitation. Occaecat elit amet commodo Lorem id aliquip do excepteur commodo sint reprehenderit aliqua nulla quis. Laboris enim nostrud sunt eu exercitation ad. Nisi tempor voluptate qui laboris id eu voluptate labore pariatur qui incididunt.",
				},
				{
					id: "002",
					votes: 1,
					title: "Test title 2",
					description:
						"loremlAute id cillum anim adipisicing labore laboris commodo consequat esse excepteur adipisicing. Tempor consequat quis aliquip aliqua ipsum et sint cupidatat consequat. Ipsum excepteur ullamco nostrud exercitation laboris sit dolor in consectetur officia.",
				},
			],
		},
		{
			categorie: "App",
			items: [
				{
					id: "001",
					votes: 2,
					title: "Test title 1",
					description:
						"loremlExercitation cupidatat ad duis laborum ex nulla laborum esse duis Lorem Lorem. Et quis irure proident non consectetur qui duis elit. Ad pariatur ut cillum qui excepteur magna consequat commodo. Voluptate culpa qui est aute irure duis mollit ullamco ad consectetur ad nulla enim.",
				},
				{
					id: "002",
					votes: 1,
					title: "Test title 2",
					description:
						"loremlAute id cillum anim adipisicing labore laboris commodo consequat esse excepteur adipisicing. Tempor consequat quis aliquip aliqua ipsum et sint cupidatat consequat. Ipsum excepteur ullamco nostrud exercitation laboris sit dolor in consectetur officia.",
				},
			],
		},
	]);

	const handleActiveCategories = (items) => {
		setActiveCategories(items);
	};

	const handleItem = (item) => {
		console.log(item);

		setActiveItem(item);
	};

	return (
		<main class="text-gray-70 w-full flex">
			<section class="flex">
				<aside class="h-full border">
					<ul class="w-52 p-2">
						<For each={categories()}>
							{(item) => (
								<li>
									<input
										type="checkbox"
										name={item.categorie}
										id={item.categorie}
										onChange={() => {
											handleActiveCategories(item.items);
										}}
									/>
									{item.categorie}
								</li>
							)}
						</For>
					</ul>
				</aside>
				<aside class="h-full border">
					<ul class="w-52 p-2">
						<Show
							when={activeCategories}
							fallback={<p>Please select an item</p>}
						>
							<For each={activeCategories()}>
								{(item) => (
									<li
										class="p-2"
										onClick={() => {
											handleItem(item);
										}}
									>
										{item.title} ({item.votes})
									</li>
								)}
							</For>
						</Show>
					</ul>
				</aside>
			</section>
			<section class="p-2 w-full border">
				<Show when={activeItem} fallback={<p>Please select an item</p>}>
					<p>{activeItem()?.description}</p>
				</Show>
			</section>
		</main>
	);
}
