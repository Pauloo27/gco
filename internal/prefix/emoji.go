package prefix

func init() {
	Packs = append(Packs, EmojiPack)
}

var EmojiPack = PrefixPack{
	Name:        "emoji",
	Separator:   " ",
	Description: "Emoji prefixes",
	Prefixes: []*Prefix{
		{
			Description: "add feature",
			Key:         "feat",
			Value:       "✨",
		},
		{
			Description: "bug fix",
			Key:         "fix",
			Value:       "🐛",
		},
		{
			Description: "add/update ci",
			Key:         "ci",
			Value:       "👷",
		},
		{
			Description: "reformat code",
			Key:         "fmt",
			Value:       "📏",
		},
		{
			Description: "add/update documentation",
			Key:         "doc",
			Value:       "📖",
		},
		{
			Description: "initial commit",
			Key:         "init",
			Value:       "🎉",
		},
		{
			Description: "delete code",
			Key:         "del",
			Value:       "🔥",
		},
		{
			Description: "move code/files",
			Key:         "mov",
			Value:       "🚚",
		},
		{
			Description: "ren code/files",
			Key:         "ren",
			Value:       "📝",
		},
		{
			Description: "add/update test",
			Key:         "test",
			Value:       "✅",
		},
		{
			Description: "add/remove/update dependency",
			Key:         "dep",
			Value:       "🔩",
		},
		{
			Description: "add tag/release",
			Key:         "tag",
			Value:       "🔖",
		},
		{
			Description: "fix typo",
			Key:         "typo",
			Value:       "✏️",
		},
		{
			Description: "merge branches/pull requests",
			Key:         "merg",
			Value:       "🔀",
		},
		{
			Description: "improve ui/ux/accessibility",
			Key:         "ux",
			Value:       "🎨",
		},
		{
			Description: "refactor",
			Key:         "ref",
			Value:       "♻️",
		},
		{
			Description: "internationalization",
			Key:         "i18n",
			Value:       "🌐",
		},
		{
			Description: "add/update .gitignore",
			Key:         "ig",
			Value:       "🙈",
		},
		{
			Description: "add/remove development related stuff",
			Key:         "dev",
			Value:       "💻",
		},
		{
			Description: "work in progress",
			Key:         "wip",
			Value:       "🚧",
		},
	},
}
