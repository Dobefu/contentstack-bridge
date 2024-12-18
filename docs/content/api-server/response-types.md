+++
title = "Response types"
type = "default"
description = "A list of possible response types"
+++

## RoutableEntryResponse

```typescript
{
  data: { // Will be null if there's an error
    // An array of alternative locales will always
    // be returned alongside the entry itself.
    alt_locales: {
      uid: string,
      content_type: string,
      locale: string,
      slug: string,
      url: string,
    }[],

    // The entry is directly queried from Contentstack.
    entry: {
      ACL: unknown,
      _in_progress: boolean,
      _version: number,
      created_at: string, // Timestamp string
      created_by: string, // UID of the creator
      locale: string,
      parent: {
        _content_type_uid: string,
        uid: string,
      }[],
      publish_details: {
        environment: string, // The environment UID
        locale: string,
        time: string, // Timestamp string
        user: string, // The user UID
      },
      tags: string[],
      title: string,
      uid: string,
      updated_at: string, // Timestamp string
      updated_by: string, // The user UID
      url: string,
      // additional options for any other fields
    },
  },
  error?: string // Will be null unless there's an error
}
```

## ContentTypeResponse

```typescript
{
  data: { // Will be null if there's an error
    content_types: {
      DEFAULT_ACL: unknown,
      SYS_ACL: unknown,
      _version: number,
      abilities: {
        create_object: boolean,
        delete_all_objects: boolean,
        delete_object: boolean,
        get_all_objects: boolean,
        get_one_object: boolean,
        update_object: boolean,
      },
      created_at: string, // Timestamp string
      description: string,
      inbuilt_class: boolean,
      last_activity: unknown,
      maintain_revisions: boolean,
      options: {
        is_page: boolean,
        publishable: boolean,
        singleton: boolean, // Whether or not the content type supports multiple entries
        sub_title: string[],
        title: string,
        url_pattern: string,
        url_prefix: string,
        // additional options for any other fields
      },
      schema: [
        {
          data_type: string,
          display_name: string,
          field_metadata: {
            _default: boolean,
            version: number,
          },
          mandatory: boolean,
          multiple: boolean,
          non_localizable: boolean,
          uid: string,
          unique: boolean,
        }[]
      ],
      title: string,
      uid: string,
      updated_at: string, // Timestamp string
    }[],
  }
  error?: string // Will be null unless there's an error
}
```
