package db

import (
	"github.com/emilgelman/gocrud/pkg/domain"
	"reflect"
	"testing"
)

func TestMemory_Empty(t *testing.T) {
	mem:= NewMemory()
	articles:= mem.GetAll()
	if len(*articles) > 0 {
		t.Error("should be empty")
	}
}

func TestMemory_GetAll(t *testing.T) {
	mem:= NewMemory()
	article:= domain.Article{Id: "1", Title: "title", Content: "content"}
	mem.Create("1",article)
	result:= mem.GetAll()
	if res:= reflect.DeepEqual(*result, []domain.Article{article}); !res {
		t.Errorf("expected %v got %v", []domain.Article{article}, *result)
	}

}

func TestMemory_Get(t *testing.T) {
	mem:= NewMemory()
	article:= domain.Article{Id: "1", Title: "title", Content: "content"}
	mem.Create("1",article)
	result, _:= mem.Get("1")
	if res:= reflect.DeepEqual(*result, article); !res {
		t.Errorf("expected %v got %v", article, *result)
	}
	_, err:= mem.Get("2")
	if err == nil {
		t.Errorf("article with id %s should not be found", "2")
	}
}

func TestMemory_Delete(t *testing.T) {
	mem:= NewMemory()
	mem.Create("1", domain.Article{})
	if err:= mem.Delete("1"); err != nil {
		t.Error("delete should not fail")
	}
	if _, err:= mem.Get("1"); err == nil {
		t.Error("should fail")
	}
	if err := mem.Delete("2"); err == nil {
		t.Error("deletion of non existing id 2 should fail")
	}
}


func TestMemory_Update(t *testing.T) {
	mem:= NewMemory()
	article:= domain.Article{Id: "1", Title: "title", Content: "content"}
	mem.Create("1",article)

	if err:= mem.Update("2", domain.Article{}); err == nil {
		t.Error("updating non existing article should fail")
	}
	if err:= mem.Update("1", domain.Article{Id: "1", Title: "newTitle"}); err != nil {
		t.Error("update should succeed")
	}
	updatedArticle, _:= mem.Get("1")
	if updatedArticle.Title != "newTitle" {
		t.Error("id wasn't updated successfully")
	}
}